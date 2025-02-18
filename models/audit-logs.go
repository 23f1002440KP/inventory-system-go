package models

type AuditLog struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	User_id    int    `gorm:"not null" json:"user_id"`
	Action     string `gorm:"size:255;not null" json:"action"`
	TableName  string `gorm:"size:255;not null" json:"table_name"`
	RecordID   int    `gorm:"not null" json:"record_id"`
	Created_at string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	User       User   `gorm:"foreignKey:User_id;References:ID"`
}
