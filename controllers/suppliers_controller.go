package controllers

import (
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/middleware"
	"23f1002440KP/inventory-system/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type NewSupplier struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

//HandleCreateSupplierPOST handles the create supplier request
func HandleCreateSupplierPOST(database *db.SQLDB) http.Handler {
	
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					var newSupplier NewSupplier
					err := json.NewDecoder(r.Body).Decode(&newSupplier)

					if err!=nil {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
						return
					}
					
					createErr :=  database.CreateSupplier(&db.SupplierStruct{
						Name:   newSupplier.Name,
						Email:  newSupplier.Email,
						Phone:  newSupplier.PhoneNumber,
						Address: newSupplier.Address,
					})

					if createErr != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusCreated, map[string]string{
						"message": "Supplier created successfully",
						"name": newSupplier.Name,
						"email": newSupplier.Email,
						"phone_number": newSupplier.PhoneNumber,
						"address": newSupplier.Address,
				})


				})))
}

//HandleGetAllSuppliersGET handles the get all suppliers request
func HandleGetAllSuppliersGET(database *db.SQLDB) http.Handler{
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					suppliers,err := database.GetAllSuppliers()
					
					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}
					utils.WriteJson(w, http.StatusOK, suppliers)


				})))
}

//HandleGetSupplierByIDGET handles the get supplier by ID request
func HandleGetSupplierByIdGET(database *db.SQLDB) http.Handler{
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {

					supplierId := strings.TrimPrefix(r.URL.Path, "/suppliers/")
					
					if !IsNumeric(supplierId) {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("supplier id must be a number")))
						return
					}

					supplier, err := database.GetSupplierByID(supplierId)

					if err != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusOK, supplier)
				})))
}

//HandleUpdateSupplierByIDPUT handles the update supplier by ID request
func HandleUpdateSupplierByIdPUT(database *db.SQLDB) http.Handler{
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					supplierId := strings.TrimPrefix(r.URL.Path, "/suppliers/")
					if !IsNumeric(supplierId) {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("supplier id must be a number")))
						return
					}

					var updatedSupplier NewSupplier

					err := json.NewDecoder(r.Body).Decode(&updatedSupplier)

					if err != nil {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
						return
					}

					updateErr := database.UpdateSupplierByID(supplierId, &db.SupplierStruct{
						Name: updatedSupplier.Name,
						Email: updatedSupplier.Email,
						Phone: updatedSupplier.PhoneNumber,
						Address: updatedSupplier.Address,
					})

					if updateErr != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(err))
						return
					}

					utils.WriteJson(w, http.StatusOK, map[string]string{
						"message": "Supplier updated successfully",
						"name": updatedSupplier.Name,
						"email": updatedSupplier.Email,
						"phone_number": updatedSupplier.PhoneNumber,
						"address": updatedSupplier.Address,
					})

				})))
}

//HandleDeleteSupplierByIDDELETE handles the delete supplier by ID request
func HandleDeleteSupplierByIdDELETE(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.JWTMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {

					supplierId := strings.TrimPrefix(r.URL.Path,"/suppliers/")
					
					if !IsNumeric(supplierId) {
						utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("supplier id must be a number")))
						return
					}

					deleteErr := database.DeleteSupplierByID(supplierId)

					if deleteErr != nil {
						utils.WriteJson(w, http.StatusInternalServerError, utils.GeneralError(deleteErr))
						return
					}

					utils.WriteJson(w, http.StatusOK, map[string]string{
						"message": "Supplier deleted successfully",
					})


				})))
}