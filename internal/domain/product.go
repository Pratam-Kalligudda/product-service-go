package domain

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"PrimaryKey; index;"`
	Name        string    `json:"name" gorm:"not null; index"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id" gorm:"index"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerId    uint      `json:"seller_id" gorm:"index"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// func NewProduct(name, description string, catId, sellerId uint, price float64, stock int) Product {
// 	return Product{Name: name, Description: description, CategoryID: catId, SellerId: sellerId, Price: price, Stock: stock}
// }

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey; index"`
	Name      string    `json:"name" gorm:"not null; unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
