package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/i101dev/rss-aggregator/controllers"
)

func NewRouter() *chi.Mux {

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

	v1Router.Get("/test", controllers.HandleTest)
	v1Router.Get("/error", controllers.HandleError)

	v1Router.Post("/users", controllers.CreateUser)
	v1Router.Get("/users", controllers.GetAllUsers)
	v1Router.Get("/users/{id}", controllers.GetUserByID)
	v1Router.Put("/users/{id}", controllers.UpdateUser)
	v1Router.Delete("/users/{id}", controllers.DeleteUser)

	router.Mount("/v1", v1Router)

	return router
}
