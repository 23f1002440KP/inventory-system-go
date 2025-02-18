package models

type Type string

const (
	StockIn  Type = "stock_in"
	StockOut Type = "stock_out"
)

type Transaction struct {
	ID         int     `gorm:"primary_key;auto_increment" json:"id"`
	Product_id int     `gorm:"not null" json:"product_id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	Type       Type    `gorm:"not null" json:"type"`
	Reason     string  `gorm:"size:255" json:"reason"`
	Created_by int     `gorm:"not null" json:"created_by"`
	Created_at string  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Product    Product `gorm:"foreignKey:Product_id;References:ID"`
	User       User    `gorm:"foreignKey:Created_by;References:ID"`
}
