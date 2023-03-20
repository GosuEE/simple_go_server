package router

import (
	"net/http"

	"github.com/rs/cors"
)

func Router(){
	serveMux := http.NewServeMux()

	UserRouter(serveMux)

    handler := cors.Default().Handler(serveMux)

    corHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           0,
		Debug:            true,
    })

    handler = corHandler.Handler(handler)

    http.ListenAndServe(":3001", handler)
}