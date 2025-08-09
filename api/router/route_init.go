package router

import (
	"net/http"
	"os"

	"github.com/LuizFernando991/golang-api/api/middlewares"
)

func InitializeRouter() http.Handler {
	r := NewRouter()

	r.Handle("GET", "/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health check called"))
	})

	var handler http.Handler = r

	if os.Getenv("ENVIRONMENT") != "DEVELOPMENT" {
		handler = middlewares.LoggerMiddleware(handler)
	}

	handler = middlewares.JSONContentTypeMiddleware([]string{"/health"})(handler)
	handler = middlewares.JSONResponseMiddleware(handler)

	controllers := GetControllers()

	InitRoutes(controllers, r)

	return handler
}
