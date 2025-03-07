package routes

import (
	db "23f1002440KP/inventory-system/database"
	"net/http"
)

type RoutesType struct{
	Router *http.ServeMux
	Database *db.SQLDB
}

func AllRoutes(routesType RoutesType) {
	routesType.Router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Inventory System"))
	}) // index Route

	
	AuthRoutes(routesType.Router, routesType.Database) // Authentication Routes 
	CategoryRoutes(routesType.Router, routesType.Database) // Category Routes

}