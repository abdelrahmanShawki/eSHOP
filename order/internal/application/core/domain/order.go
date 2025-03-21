package domain

import (
	"time"
)

type OrderItem struct {
	ProductCode string  `json:"product_code"` // Unique identifier for the product
	UnitPrice   float32 `json:"unit_price"`   // Price per unit of the product
	Quantity    int32   `json:"quantity"`     // Quantity of the product in the order
}

type Order struct {
	ID         int64       `json:"id"`          // Unique order identifier
	CustomerID int64       `json:"customer_id"` // ID of the customer placing the order
	Status     string      `json:"status"`      // Current order status (e.g., "Pending", "Shipped")
	OrderItems []OrderItem `json:"order_items"` // List of items in the order
	CreatedAt  int64       `json:"created_at"`  // Timestamp when the order was created
}

func NewOrder(customerID int64, orderItems []OrderItem) Order {
	return Order{
		CustomerID: customerID,        // Assign the provided customer ID
		Status:     "Pending",         // Default status for a new order
		OrderItems: orderItems,        // Assign the provided order items
		CreatedAt:  time.Now().Unix(), // Set the creation timestamp
	}
}
