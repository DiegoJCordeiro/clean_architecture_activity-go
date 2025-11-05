package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	ErrInvalidID         = errors.New("invalid order ID")
	ErrInvalidCustomerID = errors.New("invalid customer ID")
	ErrInvalidPrice      = errors.New("price must be greater than zero")
	ErrInvalidTax        = errors.New("tax must be greater than or equal to zero")
)

// Order representa uma ordem no sistema
type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID string             `bson:"customer_id" json:"customer_id"`
	Price      float64            `bson:"price" json:"price"`
	Tax        float64            `bson:"tax" json:"tax"`
	FinalPrice float64            `bson:"final_price" json:"final_price"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

// NewOrder cria uma nova ordem validada
func NewOrder(customerID string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:         primitive.NewObjectID(),
		CustomerID: customerID,
		Price:      price,
		Tax:        tax,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := order.Validate(); err != nil {
		return nil, err
	}

	order.CalculateFinalPrice()
	return order, nil
}

// Validate valida os dados da ordem
func (o *Order) Validate() error {
	if o.CustomerID == "" {
		return ErrInvalidCustomerID
	}
	if o.Price <= 0 {
		return ErrInvalidPrice
	}
	if o.Tax < 0 {
		return ErrInvalidTax
	}
	return nil
}

// CalculateFinalPrice calcula o preço final com taxa
func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
