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

	log.Println("API is running on port 8080")
	http.ListenAndServe(":8080", router)
}
