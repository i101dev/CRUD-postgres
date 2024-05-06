package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/i101dev/rss-aggregator/handlers"
)

func BuildRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"Get", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/test", handlers.HandleTest)
	v1Router.Get("/error", handlers.HandleError)

	router.Mount("/v1", v1Router)
	return router
}
