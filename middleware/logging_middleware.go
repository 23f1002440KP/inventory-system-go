package middleware

import (
	"23f1002440KP/inventory-system/logger"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger().Warn("Handling Request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w,r)
	})
}

