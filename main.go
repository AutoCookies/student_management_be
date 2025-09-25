package main

import (
	"log"
	"net/http"
	"student_management_be/controllers"
	"student_management_be/libs"
	"student_management_be/repository"
)

func main() {
	// Kết nối MySQL
	if err := libs.ConnectMySQL("root", "Quanlatui777****", "localhost:3306", "student_management"); err != nil {
		log.Fatal("Failed to connect MySQL:", err)
	}

	// Repository
	accountRepo := repository.NewMySQLAccountRepository(libs.DB)

	// Controller
	accountController := controllers.NewAccountController(accountRepo)

	// Router
	mux := http.NewServeMux()
	mux.Handle("/accounts", accountController)
	mux.Handle("/accounts/", accountController)

	// Run server
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
