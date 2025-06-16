package domain

import "time"

type Product struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CategoryID  uint      `json:"category_id,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Stock       int       `json:"stock,omitempty"`
	SellerId    uint      `json:"seller_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Category struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
