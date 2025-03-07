package controllers

import (
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/middleware"
	"23f1002440KP/inventory-system/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type NewCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// HandleCreateCategoryPOST handles the create category request
func HandleCreateCategoryPOST(databse *db.SQLDB) http.Handler {
	// handle the create category request
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// handle the create category request
					var newCategory NewCategory
					err := json.NewDecoder(r.Body).Decode(&newCategory)

					if errors.Is(err, io.EOF) {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("request body is empty")))
						return
					}

					if err != nil {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
						return
					}

					err = databse.CreateCategory(&db.CreateCategory{
						Name:        newCategory.Name,
						Description: newCategory.Description,
					})

					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusCreated, map[string]string{
						"message":     "Category created successfully",
						"name":        newCategory.Name,
						"description": newCategory.Description,
					})

				})))
}

// HandleGetAllCategoriesGET handles the get all categories request
func HandleGetAllCategoriesGET(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// handle the get all categories request
					categories, err := database.GetAllCategories()
					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusOK, categories)
				})))
}

// HandleGetCategoryByIDGET handles the get category by ID request
func HandleGetCategoryByIDGET(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					// Handle the get category by ID request
					categoryId := r.URL.Query().Get("id")
					category, err := database.GetCategoryByID(categoryId)
					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}
					utils.WriteJson(w, http.StatusOK, category)
				})))
}

// HandleUpdateCategoryByIDPUT handles the update category by ID request
func HandleUpdateCategoryByIDPUT(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					//handle the update category by ID request

					categoryId := r.URL.Query().Get("id")
					var updatedCategory NewCategory
					err := json.NewDecoder(r.Body).Decode(&updatedCategory)

					if errors.Is(err, io.EOF) {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("request body is empty")))
						return
					}

					if err != nil {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
						return
					}

					err = database.UpdateCategoryByID(categoryId, &db.CreateCategory{
						Name:        updatedCategory.Name,
						Description: updatedCategory.Description,
					})
					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusOK, map[string]string{
						"message":     "Category updated successfully",
						"name":        updatedCategory.Name,
						"description": updatedCategory.Description,
					})

				})))
}

// HandleDeleteCategoryByIDDELETE handles the delete category by ID request

func HandleDeleteCategoryByIDDELETE(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					//handle the delete category by ID request
					categoryId := r.URL.Query().Get("id")
					err := database.DeleteCategoryByID(categoryId)
					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusOK, map[string]string{
						"message": "Category deleted successfully",
					})
				})))
}
