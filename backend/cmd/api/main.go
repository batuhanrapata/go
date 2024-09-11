package main

import (
	"backend/config"
	service "backend/internal/casestudy"
	"backend/pkg/casestudy"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()

	// Veritabanı bağlantısını al
	dbConn := config.GetDBConn()

	// GORM ile PostgreSQL bağlantısı
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Veritabanı migrasyonu
	db.AutoMigrate(&casestudy.CaseStudy{})

	// Router ayarı ve endpoint'ler
	router := service.NewRouter(db)

	log.Println("API is running on port 8080")
	http.ListenAndServe(":8080", router)
}
