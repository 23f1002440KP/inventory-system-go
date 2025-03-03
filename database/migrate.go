package db

import (
	"23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/models"

)

func (db *SQLDB) AutoMigration() error {
	err := db.Db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Supplier{},
		&models.User{},
		&models.AuditLog{},
		&models.Inventory{},
		&models.Order{},
		&models.OrderItem{},
		&models.Transaction{},
	)
	if err != nil {
		logger.Logger().Error("Not able to migrate the database:","err",err)
		return err
	}
	logger.Logger().Info("Database migration successful")
	return nil
}