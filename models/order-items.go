package models

type OrderItem struct {
	ID         int     `gorm:"primary_key;auto_increment" json:"id"`
	Order_id   int     `gorm:"not null" json:"order_id"`
	Product_id int     `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	Price      float64 `gorm:"not null" json:"price"`
	Created_at string  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Order      Order   `gorm:"foreignKey:Order_id;References:ID"`
	Product    Product `gorm:"foreignKey:Product_id;References:ID"`
}
