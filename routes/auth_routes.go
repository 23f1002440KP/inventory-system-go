package routes

import (
	"23f1002440KP/inventory-system/controllers"
	db "23f1002440KP/inventory-system/database"
	"net/http"
)

func AuthRoutes(router *http.ServeMux, database *db.SQLDB) {

	router.Handle("POST /register",controllers.HandleRegisterPOST(database))
	router.Handle("POST /login",controllers.HandleLoginPOST(database))
}
