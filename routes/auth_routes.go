package routes

import (
	"23f1002440KP/inventory-system/controllers"
	db "23f1002440KP/inventory-system/database"
	"net/http"
)

func AuthRoutes(router *http.ServeMux, database *db.SQLDB) {

	router.Handle("POST /auth/register",controllers.HandleRegisterPOST(database))
	router.Handle("POST /auth/login",controllers.HandleLoginPOST(database))
	router.Handle("GET /auth/me",controllers.HandleGetUserGET(database))
}
