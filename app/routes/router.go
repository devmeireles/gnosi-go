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
	// auth.HandleFunc("/register", controller.CreateUser).Methods("POST")

	// Swagger Routes
	// r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	api := r.PathPrefix("/api").Subrouter()
	// api.Use(middlewares.AuthJwtVerify)

	// Categories routes
	api.HandleFunc("/category", controller.GetAllCategories).Methods("GET")

	// Skill routes
	// api.HandleFunc("/skills", controller.GetAllSkills).Methods("GET")
	// api.HandleFunc("/skill/{id:[0-9]+}", controller.GetSkill).Methods("GET")
	// api.HandleFunc("/skill", controller.CreateSkill).Methods("POST")
	// api.HandleFunc("/skill/{id:[0-9]+}", controller.UpdateSkill).Methods("PUT")
	// api.HandleFunc("/skill/{id:[0-9]+}", controller.DeleteSkill).Methods("DELETE")

	// User routes
	// api.HandleFunc("/user/address", controller.CreateUserAddress).Methods("POST")
	// api.HandleFunc("/user/address", controller.UpdateUserAddress).Methods("PUT")
	// api.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
	// api.HandleFunc("/user/{id:[0-9]+}", controller.GetUser).Methods("GET")
	// api.HandleFunc("/user/{id:[0-9]+}", controller.UpdateUser).Methods("PUT")
	// api.HandleFunc("/user/{id:[0-9]+}", controller.DeleteUser).Methods("DELETE")

	server.Router = r

	return r
}
