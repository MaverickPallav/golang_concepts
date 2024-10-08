package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strconv"
    "fmt"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

type Order struct {
    ID        int       `json:"id"`
    ProductID int       `json:"productId"`
    Quantity  int       `json:"quantity"`
    Total     float64   `json:"total"`
}

var products []Product
var orders []Order

// Get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

// Add a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    _ = json.NewDecoder(r.Body).Decode(&product)
    product.ID = len(products) + 1
    products = append(products, product)
    json.NewEncoder(w).Encode(product)
}

// Get product by ID
func getProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, item := range products {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Product not found", http.StatusNotFound)
}

// Update product by ID
func updateProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, item := range products {
        if item.ID == id {
            products = append(products[:index], products[index+1:]...)
            var updatedProduct Product
            _ = json.NewDecoder(r.Body).Decode(&updatedProduct)
            updatedProduct.ID = id
            products = append(products, updatedProduct)
            json.NewEncoder(w).Encode(updatedProduct)
            return
        }
    }
    http.Error(w, "Product not found", http.StatusNotFound)
}

// Delete product by ID
func deleteProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for index, item := range products {
        if item.ID == id {
            products = append(products[:index], products[index+1:]...)
            json.NewEncoder(w).Encode(products)
            return
        }
    }
}

// Get all orders
func getOrders(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

// Place a new order
func createOrder(w http.ResponseWriter, r *http.Request) {
    var order Order
    _ = json.NewDecoder(r.Body).Decode(&order)
    order.ID = len(orders) + 1
    for _, product := range products {
        if product.ID == order.ProductID {
            order.Total = product.Price * float64(order.Quantity)
        }
    }
    orders = append(orders, order)
    json.NewEncoder(w).Encode(order)
}

// Get order by ID
func getOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    for _, item := range orders {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Order not found", http.StatusNotFound)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Ecommerce API!")
}


func main() {
    router := mux.NewRouter()

    // '/' route
    router.HandleFunc("/", WelcomeHandler).Methods("GET")

    // Product routes
    router.HandleFunc("/products", getProducts).Methods("GET")
    router.HandleFunc("/products", createProduct).Methods("POST")
    router.HandleFunc("/products/{id}", getProduct).Methods("GET")
    router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
    router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

    // Order routes
    router.HandleFunc("/orders", getOrders).Methods("GET")
    router.HandleFunc("/orders", createOrder).Methods("POST")
    router.HandleFunc("/orders/{id}", getOrder).Methods("GET")

    fmt.Println("Server started on http://localhost:4001")
    log.Fatal(http.ListenAndServe(":4001", router))
}
