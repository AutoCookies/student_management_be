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
	classroomRepo := repository.NewMySQLClassroomRepository(libs.DB)

	// Controller
	accountController := controllers.NewAccountController(accountRepo)
	classroomController := controllers.NewClassroomController(classroomRepo)

	// Router
	mux := http.NewServeMux()

	// Accounts
	mux.Handle("/accounts", accountController)
	mux.Handle("/accounts/", accountController)

	// Classrooms
	mux.Handle("/classrooms", classroomController)
	mux.Handle("/classrooms/", classroomController)

	// Run server
	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
