package models

type Status string

const (
	OrderStatusPending   Status = "pending"
	OrderStatusShipped   Status = "shipped"
	OrderStatusDelivered Status = "delivered"
	OrderStatusCancelled Status = "cancelled"
)

type Order struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	User_id    int    `gorm:"not null" json:"user_id"`
	Status     Status `gorm:"not null" json:"status" default:"pending"`
	Total      int    `gorm:"not null" json:"total"`
	Created_at string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	User       User   `gorm:"foreignKey:User_id;References:ID"`
}
