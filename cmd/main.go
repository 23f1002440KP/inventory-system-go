package main

import (
	"23f1002440KP/inventory-system/configs"
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/logger"
	"fmt"

	"gorm.io/gorm"
) 


func main() {
	fmt.Println("Starting Inventory System")
	newLogger := logger.Logger()

	// Loading the configurations
	newLogger.Info("Loading the configurations")
	cfgpath := configs.LoadConfigs()


	//loading the database
	newLogger.Info("Loading the database")
	db,err := db.Connect(cfgpath)
	if err != nil {
		newLogger.Error("Not able to connect to the database:","error",err)
	}
	fmt.Println(db)
	//loading the server

	//loading the routes

	//starting the server

	//listening to the server

	//shutting down the server

}