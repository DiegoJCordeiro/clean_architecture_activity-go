package dtos

// CreateOrderInputDTO representa os dados de entrada
type CreateOrderInputDTO struct {
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
}

// CreateOrderOutputDTO representa os dados de saída
type CreateOrderOutputDTO struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
}

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
