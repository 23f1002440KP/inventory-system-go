package models

type Inventory struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	Product_id  int    `gorm:"not null" json:"product_id"`
	Stock       int    `gorm:"not null;default:0" json:"stock"`
	Updated_at  string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Product     Product `gorm:"foreignKey:Product_id;References:ID"`
}
