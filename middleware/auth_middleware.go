package middleware

import (
	"23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/utils"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)



func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the request has a valid jwt token
		tokenString := r.Header.Get("Authorization")

		token,err := utils.ValidateJWT(tokenString)
		if token == nil || err != nil {
			utils.WriteJson(w,http.StatusUnauthorized,utils.GeneralError(err))
			logger.Logger().Error("Unauthorized access","token",tokenString,"error",err)
			return
		}

		user,ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			utils.WriteJson(w,http.StatusUnauthorized,utils.GeneralError(err))
			logger.Logger().Error("Unauthorized access","token",tokenString,"error","invalid token")
			return
		}

		
		userId := user["payload"].(map[string]interface{})["id"]

		ctx := context.WithValue(r.Context(),utils.UserIDKey,userId)
		next.ServeHTTP(w,r.WithContext(ctx))
		// if not return a 401 unauthorized error
		// if the token is valid, call the next handler
	})
}
