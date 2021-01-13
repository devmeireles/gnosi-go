package routes

import (
	"github.com/devmeireles/gnosi-api/app/controller"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

// SetupRoutes is comment about function
func (server *Server) SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Auth Routes
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", controller.Login).Methods("POST")
	auth.HandleFunc("/register", controller.CreateUser).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	// api.Use(middlewares.AuthJwtVerify)

	// Categories routes
	api.HandleFunc("/category", controller.GetCategories).Methods("GET")
	api.HandleFunc("/category/{id:[0-9]+}", controller.GetCategory).Methods("GET")
	api.HandleFunc("/category", controller.CreateCategory).Methods("POST")
	api.HandleFunc("/category/{id:[0-9]+}", controller.UpdateCategory).Methods("PUT")
	api.HandleFunc("/category/{id:[0-9]+}", controller.DeleteCategory).Methods("DELETE")

	// Catalogues routes
	api.HandleFunc("/catalogue", controller.GetCatalogues).Methods("GET")
	api.HandleFunc("/catalogue/{id:[0-9]+}", controller.GetCatalogue).Methods("GET")
	api.HandleFunc("/catalogue", controller.CreateCatalogue).Methods("POST")
	api.HandleFunc("/catalogue/{id:[0-9]+}", controller.UpdateCatalogue).Methods("PUT")
	api.HandleFunc("/catalogue/{id:[0-9]+}", controller.DeleteCatalogue).Methods("DELETE")

	// Seasons routes
	api.HandleFunc("/season", controller.GetSeasons).Methods("GET")
	api.HandleFunc("/season/{id:[0-9]+}", controller.GetSeason).Methods("GET")
	api.HandleFunc("/season", controller.CreateSeason).Methods("POST")
	api.HandleFunc("/season/{id:[0-9]+}", controller.UpdateSeason).Methods("PUT")
	api.HandleFunc("/season/{id:[0-9]+}", controller.DeleteSeason).Methods("DELETE")

	// api.HandleFunc("/catalogue", controller.GetAllCatalogues).Methods("GET")

	server.Router = r

	return r
}
