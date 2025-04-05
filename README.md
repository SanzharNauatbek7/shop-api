# 🛍️ Shop API

Simple REST API built with Go (Gin) and MongoDB for managing products in a shop.

## 📦 Features

- Add, update, delete products
- Get all products or get product by ID
- MongoDB integration
- Clean and modular code structure

## 🧰 Stack

- Go 1.21+ (или Go 1.24, если ты используешь)
- Gin Web Framework
- MongoDB
- Postman (для тестов)

## 🛠️ API Endpoints

| Method | Endpoint        | Description            |
|--------|------------------|------------------------|
| GET    | `/products`      | Get all products       |
| GET    | `/products/:id`  | Get product by ID      |
| POST   | `/products`      | Create a new product   |
| PUT    | `/products/:id`  | Update product by ID   |
| DELETE | `/products/:id`  | Delete product by ID   |

## 📂 MongoDB

- Database: `shopdb`
- Collection: `product`

## ▶️ Run Locally

```bash
go run main.go
