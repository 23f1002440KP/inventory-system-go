package main

import (
	"23f1002440KP/inventory-system/configs"
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

)

func main() {
	fmt.Println("Starting Inventory System")
	newLogger := logger.Logger()

	// Loading the configurations
	newLogger.Info("Loading the configurations")
	cfgpath := configs.LoadConfigs()

	//loading the database
	newLogger.Info("Loading the database")
	database, err := db.Connect(cfgpath)
	if err != nil {
		newLogger.Error("Not able to connect to the database:", "error", err)
		return 
	}
	newLogger.Info("Database connection established")
	newLogger.Info("Migrating the database")

	err_migrate := database.AutoMigration()

	if err_migrate != nil {
		newLogger.Error("Not able to migrate the database:", "error", err_migrate)
		return
	}

	//seeding the database
	database.Seed()


	//loading the routes

	router := http.NewServeMux()

	routesType := routes.RoutesType{
		Router: router,
		Database: database,
	}

	routes.AllRoutes(routesType)


	//starting the server
	addr := cfgpath.Server.Host + ":" + cfgpath.Server.Port

	server := http.Server{
		Addr: addr,
		Handler: router,
	}

	logger.Logger().Info("Starting the server at", "address", fmt.Sprintf("http://%s", addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Logger().Error("Error starting the server", "error", err)
		}
	}()


	//shutting down the server
	<- done 
	logger.Logger().Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger().Error("Error shutting down the server", "error", err)
	}
	logger.Logger().Info("Server shutdown successfully")
}
