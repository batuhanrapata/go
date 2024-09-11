package db

import (
    "log"
    "gorm.io/gorm"
    "backend/pkg/casestudy"
	
)

// MigrateTables migrasyon işlemi gerçekleştirir
func MigrateTables(db *gorm.DB) {
    // GORM ile tabloları oluşturuyoruz
    err := db.AutoMigrate(&casestudy.CaseStudy{})
    if err != nil {
        log.Fatalf("Could not migrate database: %v", err)
    }
    log.Println("Database migration completed")
}
