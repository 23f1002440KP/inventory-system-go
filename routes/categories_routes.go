package routes

import (
	"23f1002440KP/inventory-system/controllers"
	db "23f1002440KP/inventory-system/database"
	"net/http"
)

func CategoryRoutes(router *http.ServeMux, database *db.SQLDB) {
	router.Handle("POST /categories", controllers.HandleCreateCategoryPOST(database))
	router.Handle("GET /categories", controllers.HandleGetAllCategoriesGET(database))
	router.Handle("GET /categories/{id}", controllers.HandleGetCategoryByIDGET(database))
	router.Handle("PUT /categories/{id}", controllers.HandleUpdateCategoryByIDPUT(database))
	router.Handle("DELETE /categories/{id}", controllers.HandleDeleteCategoryByIDDELETE(database))

}
