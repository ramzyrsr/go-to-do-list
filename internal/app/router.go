package app

import (
	"log"
	"net/http"
	"to-do-list/internal/domain/repository"
	"to-do-list/internal/domain/service"
	"to-do-list/internal/infrastructure/db"
	"to-do-list/internal/infrastructure/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() {
	// Set up the database connection
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Set up the repositories, services, and handlers for User and Product
	userRepository := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	productRepository := repository.NewProductRepository(dbConn)
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	// Create a new router using mux
	r := mux.NewRouter()

	// Apply the Content-Type middleware globally
	// r.Use(middleware.SetJSONContentType)

	// Define routes for User
	r.HandleFunc("/sign-up", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/sign-in", userHandler.Login).Methods("POST")
	r.HandleFunc("/sign-out", userHandler.Logout).Methods("POST")

	// Define routes for Product
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")        // POST /products
	r.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")   // GET /products/{id}
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")    // PUT /products/{id}
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE") // DELETE /products/{id}

	// Start the HTTP server with the router
	http.Handle("/", r)
}
