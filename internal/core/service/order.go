package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallenge/internal/adapter/handler/dto"
	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
	"github.com/mfritschdotgo/techchallenge/internal/core/port"
)

type OrderService struct {
	orderRepo      port.OrderRepository
	clientService  *ClientService
	productService *ProductService
}

func NewOrderService(repo port.OrderRepository, clientService *ClientService, productService *ProductService) *OrderService {
	return &OrderService{
		orderRepo:      repo,
		clientService:  clientService,
		productService: productService,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, dto dto.CreateOrderRequest) (*domain.Order, error) {
	_, err := s.clientService.GetClientByCPF(ctx, dto.Client)
	if err != nil {
		return nil, fmt.Errorf("client validation failed: %w", err)
	}

	total := 0.0
	productDetails := make(map[string]struct {
		Price float64
		Name  string
	})

	for _, item := range dto.Products {
		product, err := s.productService.GetProductByID(ctx, item.ID)
		if err != nil {
			return nil, fmt.Errorf("product validation failed for product ID %s: %w", item.ID, err)
		}
		productDetails[item.ID] = struct {
			Price float64
			Name  string
		}{
			Price: product.Price,
			Name:  product.Name,
		}
		total += product.Price * float64(item.Quantity)
	}

	items := ConvertDTOtoSlice(dto.Products, productDetails)

	order, err := domain.NewOrder(dto.Client, items, 0, total, "created")
	if err != nil {
		return nil, fmt.Errorf("failed to create order instance: %w", err)
	}

	savedOrder, err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return savedOrder, nil
}

func ConvertDTOtoSlice(dtoProducts []dto.ProductItem, productDetails map[string]struct {
	Price float64
	Name  string
}) []domain.OrderItem {
	var domainItems []domain.OrderItem
	for _, item := range dtoProducts {
		details := productDetails[item.ID]
		domainItems = append(domainItems, domain.OrderItem{
			ProductID:   item.ID,
			ProductName: details.Name,
			Quantity:    item.Quantity,
			Price:       details.Price * float64(item.Quantity),
		})
	}
	return domainItems
}

func (s *OrderService) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	order, err := s.orderRepo.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	return order, nil
}

func (s *OrderService) GetOrders(ctx context.Context, page, size int) ([]domain.Order, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	orders, err := s.orderRepo.GetOrders(ctx, page, size)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderService) SetOrderStatus(ctx context.Context, id string, status int) (*domain.OrderStatus, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	_, err = s.orderRepo.GetOrderByID(ctx, uuidID.String())

	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	orderStatus, err := domain.SetStatus(status)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = s.orderRepo.SetStatus(ctx, uuidID, orderStatus.Status, orderStatus.StatusDescription)

	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %w", err)
	}

	return orderStatus, nil
}
