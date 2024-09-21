/*
 * Order Service API
 *
 * API for managing orders in the e-commerce platform
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Order struct {

	OrderId string `json:"orderId,omitempty"`

	UserId string `json:"userId,omitempty"`

	Status *OrderStatus `json:"status,omitempty"`

	OrderDate time.Time `json:"orderDate,omitempty"`

	TotalAmount float64 `json:"totalAmount,omitempty"`

	Items []OrderItem `json:"items,omitempty"`

	ShippingAddress *Address `json:"shippingAddress,omitempty"`

	BillingAddress *Address `json:"billingAddress,omitempty"`
}
