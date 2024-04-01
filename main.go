package main

import (
	"belajar-golang/app/controller"
	"belajar-golang/app/middleware"
	"belajar-golang/connection"
	"fmt"
	"net/http"
)

func main() {
	// connection
	connection.InitDB()

	// end point user
	http.HandleFunc("/users", controller.GetUsers)
	http.HandleFunc("/create", controller.CreateUserHandler)
	http.HandleFunc("/delete", controller.DeleteUserHandler)
	http.HandleFunc("/login", controller.LoginHandler)

	// end point products
	http.HandleFunc("/products", controller.GetAllProductsHandler)
	http.Handle("/products/create", middleware.JWTAuthMiddlewareProduct(http.HandlerFunc(controller.CreateProductHandler)))
	http.Handle("/products/delete", middleware.JWTAuthMiddlewareProduct(http.HandlerFunc(controller.DeleteProductHandler)))
	http.Handle("/products/update", middleware.JWTAuthMiddlewareProduct(http.HandlerFunc(controller.UpdateProductHandler)))
	http.HandleFunc("/products/id", controller.GetProductByIdHandler)

	// end point dashboard
	http.Handle("/dashboard", middleware.JWTAuthMiddleware(http.HandlerFunc(controller.DahshboardHandler)))

	// run server GO
	fmt.Println("Server is running on http://localhost:5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting on server", err.Error())
	}
}
