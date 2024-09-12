package main

import (
	"backend/config"
	service "backend/internal/casestudy"
	"backend/internal/firebase"
	"backend/pkg/casestudy"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gorilla/handlers"
)

func main() {
	config.LoadEnv()

	err := firebase.InitFirebase()
	if err != nil {
		log.Fatal("Firebase initialization failed:", err)
	}

	dbConn := config.GetDBConn()

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.AutoMigrate(&casestudy.CaseStudy{})

	router := service.NewRouter(db)

	// CORS ayarlarını yapılandır
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// CORS middleware'ini router'a ekleyin
	corsHandler := handlers.CORS(corsHeaders, corsOrigins, corsMethods)

	log.Println("API is running on port 8080")
	http.ListenAndServe(":8080", corsHandler(router))
}
