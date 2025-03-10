package models

type Supplier struct {
	ID      int    `gorm:"primary_key;auto_increment" json:"id"`
	Name    string `gorm:"size:255;not null" json:"name"`
	Email   string `gorm:"size:100;not null;unique" json:"email"`
	Phone   string `gorm:"size:100;not null;unique" json:"phone_number"`
	Address string `gorm:"size:255;not null" json:"address"`
	IsDeleted bool 	`gorm:"default:false" json:"is_deleted"`
	Created_at string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
