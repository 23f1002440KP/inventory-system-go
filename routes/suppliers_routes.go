package routes

import (
	"23f1002440KP/inventory-system/controllers"
	db "23f1002440KP/inventory-system/database"
	"net/http"
)

func SupplierRoutes(router *http.ServeMux, database *db.SQLDB) {
	
	router.Handle("POST /suppliers",controllers.HandleCreateSupplierPOST(database))
	router.Handle("GET /suppliers",controllers.HandleGetAllSuppliersGET(database))
	router.Handle("GET /suppliers/{id}",controllers.HandleGetSupplierByIdGET(database))
	router.Handle("PUT /suppliers/{id}",controllers.HandleUpdateSupplierByIdPUT(database))
	router.Handle("DELETE /suppliers/{id}",controllers.HandleDeleteSupplierByIdDELETE(database))

}