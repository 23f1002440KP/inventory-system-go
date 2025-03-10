package models

type Category struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
	IsDeleted  bool  `gorm:"default:false" json:"is_deleted"`
}
