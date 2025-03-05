package controllers

import (
	db "23f1002440KP/inventory-system/database"
	"23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/middleware"
	"23f1002440KP/inventory-system/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator"
)
type RegisterRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role string `json:"role" validate:"required,oneof=admin staff"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func  HandleRegisterPOST(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// handle the register request
		
				var user RegisterRequest
		
				err := json.NewDecoder(r.Body).Decode(&user)
				
				if errors.Is(err,io.EOF){
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("request body is empty")))
					return
				}
				
				if err != nil {
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
					return	
				}
				
		
				if err := validator.New().Struct(user); err != nil {
		
					ValidateErrs := err.(validator.ValidationErrors)
					utils.WriteJson(w, http.StatusBadRequest, utils.ValidationError(ValidateErrs))
					return
				}
		
		
				lastId,err := database.Register(
					user.Name,
					user.Email,
					user.Password,
					string(user.Role),
				)
		
				if err !=nil {
					logger.Logger().Error("Error Registering User","error",err)
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
				}else{
					logger.Logger().Info("User Registered Successfully")
					utils.WriteJson(w, http.StatusCreated, map[string]string{"Success": "User Created","id":fmt.Sprintf("%d",lastId)})
				}
		
		
			}))
}


func HandleLoginPOST(database *db.SQLDB) http.Handler {
	return middleware.LoggingMiddleware(
		http.HandlerFunc(
			func (w http.ResponseWriter, r *http.Request) {
				// handle the login request
		
				var user LoginRequest
		
				err := json.NewDecoder(r.Body).Decode(&user)
		
				if errors.Is(err,io.EOF){
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(errors.New("request body is empty")))
					return
				}
		
				if err != nil {
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
					return	
				}
		
				if err := validator.New().Struct(user); err != nil {
					
					ValidateErrs := err.(validator.ValidationErrors)
					utils.WriteJson(w, http.StatusBadRequest, utils.ValidationError(ValidateErrs))
					return
				}
		
				userId,err := database.Login(user.Email,user.Password)
		
				if err !=nil {
					logger.Logger().Error("Error Logging In User","error",err)
					utils.WriteJson(w, http.StatusBadRequest, utils.GeneralError(err))
				}else{
					logger.Logger().Info("User Logged In Successfully")
					utils.WriteJson(w, http.StatusOK, map[string]string{"Success": "User Logged In","id":fmt.Sprintf("%d",userId)})
				}
		
			}))
	
}