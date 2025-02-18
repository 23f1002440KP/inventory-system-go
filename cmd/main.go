package main

import (
	"23f1002440KP/inventory-system/configs"
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/logger"
	"fmt"
	"strconv"

	"gorm.io/gorm"
) 

type Product struct {
	gorm.Model
	ID int
	Name string
	Price float64
	Quantity int
}

func main() {
	fmt.Println("Starting Inventory System")
	newLogger := logger.Logger()

	// Loading the configurations
	newLogger.Info("Loading the configurations")
	cfgpath := configs.LoadConfigs()
	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/inventory", cfgpath.Database.User, cfgpath.Database.Password, cfgpath.Database.Host, strconv.Itoa(cfgpath.Database.Port))

	newLogger.Info("Database connection string ","db_url", db_url)


	//loading the database
	newLogger.Info("Loading the database")
	db := db.Connect(cfgpath)
	newLogger.Info("Database connection established")
	
	//loading the server

	//loading the routes

	//starting the server

	//listening to the server

	//shutting down the server

}