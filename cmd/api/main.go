package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mfritschdotgo/techchallenge/configs"
	_ "github.com/mfritschdotgo/techchallenge/docs"
	"github.com/mfritschdotgo/techchallenge/internal/adapter/handler/httpserver"
	"github.com/mfritschdotgo/techchallenge/internal/adapter/repository"
	"github.com/mfritschdotgo/techchallenge/internal/core/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title			Skina Lanches Management API
// @version		1.0
// @description	APIs for using the management system and sales orders
// @BasePath					/
func main() {
	config := configs.GetConfig()

	client, err := connectDatabase(config.MONGO_USER, config.MONGO_PASSWORD, config.MONGO_HOST, config.MONGO_PORT, config.MONGO_DATABASE)
	if err != nil {
		panic(err)
	}
	db := client.Database(config.MONGO_DATABASE)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := httpserver.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryService)
	productHandler := httpserver.NewProductHandler(productService)

	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := httpserver.NewClientHandler(clientService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, clientService, productService)
	orderHandler := httpserver.NewOrderHandler(orderService)

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Get("/{id}", productHandler.GetProductByID)
		r.Get("/", productHandler.GetProducts)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Post("/", categoryHandler.CreateCategory)
		r.Put("/{id}", categoryHandler.UpdateCategory)
		r.Get("/{id}", categoryHandler.GetCategoryByID)
		r.Get("/", categoryHandler.GetCategories)
		r.Delete("/{id}", categoryHandler.DeleteCategory)
	})

	r.Route("/clients", func(r chi.Router) {
		r.Post("/", clientHandler.CreateClient)
		r.Get("/{cpf}", clientHandler.GetClientByCPF)
	})

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", orderHandler.GetOrders)
		r.Get("/{id}", orderHandler.GetOrderByID)
		r.Post("/", orderHandler.CreateOrder)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/doc.json")))

	http.ListenAndServe(":9090", r)
}

func connectDatabase(user string, password string, host string, port string, dbname string) (*mongo.Client, error) {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s", user, password, host, port, dbname, user)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
