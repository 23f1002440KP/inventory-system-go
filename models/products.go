package models

type Product struct {
	ID          int      `gorm:"primary_key;auto_increment" json:"id"`
	Name        string   `gorm:"size:255;not null" json:"name"`
	Description string   `gorm:"size:255" json:"description"`
	Price       float64  `gorm:"not null" json:"price"`
	Stock       int      `gorm:"not null;default:0" json:"stock"`
	Category_id int      `gorm:"not null" json:"category_id"` 
	Supplier_id int      `gorm:"not null" json:"supplier_id"` 
	Created_at  string   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at  string   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Category    Category `gorm:"foreignKey:Category_id;References:ID"`
	Supplier    Supplier `gorm:"foreignKey:Supplier_id;References:ID"`
}
