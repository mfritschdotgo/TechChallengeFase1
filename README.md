# Skina Lanches Management API

## Overview

The Skina Lanches Management API is designed to facilitate the management of products, categories, clients, and orders for a sandwich shop. The API provides endpoints to create, read, update, and delete entities within the system.
This project follows a hexagonal architecture and is based on a DDD (Domain driven design) generated from event storming in which the steps can be analyzed in:

## Table of Contents

- [Skina Lanches Management API](#skina-lanches-management-api)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Setup](#setup)
    - [Docker Setup](#docker-setup)
    - [Compose Setup](#compose-setup)
  - [API Endpoints](#api-endpoints)
    - [Categories](#categories)
    - [Clients](#clients)
    - [Orders](#orders)
    - [Products](#products)
    - [FakeCheckout](#fakeCheckout)

## Setup

### Docker Setup

To set up the application using Docker, follow these steps:

1. Ensure you have Docker installed on your machine and be logged into a registry (optional).
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Build the Docker image using the provided Dockerfile:
   ```sh
   docker build -t yourRepository/skina-lanches-api:version .
5. Run the Docker container:
    ```sh
    docker run -p 9090:9090 -e MONGO_USER=user -e MONGO_PASSWORD=password -e MONGO_PORT=port -e MONGO_HOST=localhost -e MONGO_DATABASE=database -e SWAG_HOST=localhost yourRepository/skina-lanches-api:version
6. Uploading the image to a registry:
   ```sh
    docker push yourRepository/skina-lanches-api:version
   
### Compose Setup

To set up the application using Docker compose, follow these steps:

1. Ensure you have Docker engine and docker compose installed on your machine.
2. Navigate to the compose directory.
3. Edit the .env file, setting the environment variables as desired
4. Start service from docker compose file, must be in the compose directory:
   ```sh
   docker compose up -d

## Integrated testing via Swagger  

1. Access the skinaapis service address with the port defined in the .env file:
   http://127.0.0.1:9090/docs/index.html#/
2. Add a client via the using method post endpoint [Clients](#clients).
3. Add a product via the using method post endpoint [Products](#products)
4. Add a request via the using method post endpoint [Orders](#orders)
5. Simulate a fake checkout through the endpoint using method post [FakeCheckout](#fakeCheckout)

## API Endpoints

### Categories

- **GET /categories**
  - Retrieves a paginated list of categories.
  - Parameters:
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of categories per page.
  - Responses:
    - `200`: Successfully retrieved list of categories.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /categories**
  - Adds a new category to the database.
  - Body: `dto.CreateCategoryRequest`
  - Responses:
    - `201`: Successfully created category.
    - `400`: Bad request if the category data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /categories/{id}**
  - Retrieves details of a category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Responses:
    - `200`: Successfully retrieved the category details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Category not found if the ID does not match any category.
    - `500`: Internal server error if there is a problem on the server side.

- **PUT /categories/{id}**
  - Replaced category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Body: `dto.CreateCategoryRequest`
  - Responses:
    - `200`: Successfully updated category.
    - `400`: Invalid input, object is invalid.
    - `404`: Category not found.
    - `500`: Internal server error.

- **DELETE /categories/{id}**
  - Deletes a category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Responses:
    - `200`: Message indicating successful deletion.
    - `400`: Bad request if the ID is not provided or is invalid.
    - `404`: Category not found if the ID does not match any category.
    - `500`: Internal server error if there is a problem deleting the category.
- **Patch /categories/{id}**
  - Update a category by its ID.
  - Parameters:
    - `id` (string): Category ID.
  - Responses:
    - `200`: Message indicating successful deletion.
    - `400`: Bad request if the ID is not provided or is invalid.
    - `404`: Category not found if the ID does not match any category.
    - `500`: Internal server error if there is a problem deleting the category.

### Clients

- **POST /clients**
  - Adds a new client to the database.
  - Body: `dto.CreateClientRequest`
  - Responses:
    - `201`: Client successfully created.
    - `400`: Bad request if the client data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /clients/{cpf}**
  - Retrieves details of a client by its CPF.
  - Parameters:
    - `cpf` (string): Client CPF.
  - Responses:
    - `200`: Successfully retrieved the client details.
    - `400`: Bad request if the CPF is not provided or invalid.
    - `404`: Client not found if the CPF does not match any client.
    - `500`: Internal server error if there is a problem on the server side.

### Orders

- **GET /orders**
  - Retrieves a paginated list of orders.
  - Parameters:
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of orders per page.
  - Responses:
    - `200`: Successfully retrieved list of orders.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /orders**
  - Adds a new order to the database.
  - Body: `dto.CreateOrderRequest`
  - Responses:
    - `201`: Successfully created order.
    - `400`: Bad request if the order data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /orders/{id}**
  - Retrieves details of an order by its ID.
  - Parameters:
    - `id` (string): Order ID.
  - Responses:
    - `200`: Successfully retrieved the order details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Order not found if the ID does not match any order.
    - `500`: Internal server error if there is a problem on the server side.
- **PATCH /orders/{id}/{status}**
  - Update the status of an order 
  - Parameters:
    - `id` (string): Order ID.
    - `status` (string): status.
  - Responses:
    - `200`: Successfully status updated.
    - `400`: Bad request if the Status is not provided or invalid
    - `500`: Internal server error if there is a problem on the server side.

### Products

- **GET /products**
  - Retrieves a paginated list of products, optionally filtered by category.
  - Parameters:
    - `category` (string, optional): Filter by category ID.
    - `page` (integer, default: 1): Page number for pagination.
    - `pageSize` (integer, default: 10): Number of products per page.
  - Responses:
    - `200`: Successfully retrieved list of products.
    - `500`: Internal server error if there is a problem on the server side.

- **POST /products**
  - Adds a new product to the database.
  - Body: `dto.CreateProductRequest`
  - Responses:
    - `201`: Product successfully created.
    - `400`: Bad request if the product data is invalid.
    - `500`: Internal server error if there is a problem on the server side.

- **GET /products/{id}**
  - Retrieves details of a product by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Responses:
    - `200`: Successfully retrieved the product details.
    - `400`: Bad request if the ID is not provided or invalid.
    - `404`: Product not found if the ID does not match any product.
    - `500`: Internal server error if there is a problem on the server side.

- **PUT /products/{id}**
  - Replaced product by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Body: `dto.CreateProductRequest`
  - Responses:
    - `200`: Product successfully updated.
    - `400`: Invalid input, object is invalid.
    - `404`: Product not found.
    - `500`: Internal server error.
- **PATCH /products/{id}**
  - Updates product details by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Body: `dto.CreateProductRequest`
  - Responses:
    - `200`: Product successfully updated.
    - `400`: Invalid input, object is invalid.
    - `404`: Product not found.
    - `500`: Internal server error.

- **DELETE /products/{id}**
  - Deletes a product by its ID.
  - Parameters:
    - `id` (string): Product ID.
  - Responses:
    - `200`: Message indicating successful deletion.
    - `400`: Bad request if the ID is not provided or is invalid.
    - `404`: Product not found if the ID does not match any product.
    - `500`: Internal server error if there is a problem deleting the product.# techchallenge
      
### fakeCheckout

- **POST /fakeCheckout/{id}**
  - Adds a new client to the database.
  - Parameters:
    - `id` (string): Product ID.
  - Body: `dto.CreateClientRequest`
  - Responses:
    - `200`: Successfully fake checkout.
    - `400`: Bad request if the ID is not provided or invalid.
    - `500`: Internal server error if there is a problem on the server side.
