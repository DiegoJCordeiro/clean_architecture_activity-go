package graphqls

type Order struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type CreateOrderInput struct {
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
}
