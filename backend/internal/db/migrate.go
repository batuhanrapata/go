package db

import (
    "log"
    "gorm.io/gorm"
    "backend/pkg/casestudy"
	
)

func MigrateTables(db *gorm.DB) {
    err := db.AutoMigrate(&casestudy.CaseStudy{})
    if err != nil {
        log.Fatalf("Could not migrate database: %v", err)
    }
    log.Println("Database migration completed")
}
