package db

import (
	"23f1002440KP/inventory-system/configs"
	"23f1002440KP/inventory-system/logger"
	"fmt"
	"strconv"

	"github.com/lmittmann/tint"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfgDb *configs.Configs) (*gorm.DB,error) {
	// Connect to the database
	
	if cfgDb.Database.Name == "postgres" {
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/inventory", cfgDb.Database.User, cfgDb.Database.Password, cfgDb.Database.Host, strconv.Itoa(cfgDb.Database.Port))
	logger.Logger().Debug("Connecting to the database:","db_url",db_url)
	
	db,err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		logger.Logger().Error("Not able to connect to the database:",tint.Err(err))
	} 
	logger.Logger().Info("Database connection established")
	return db,nil
	}
	
	logger.Logger().Error("Database is not supported","db_name",cfgDb.Database.Name)
	return nil,fmt.Errorf("Database is not supported")
	
}