package main

import (
	"backend/config"
	"backend/internal/casestudy"
	transport "backend/internal/casestudy/transport"
	"backend/internal/firebase"
	model "backend/pkg/casestudy"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "caller", log.DefaultCaller)

	config.LoadEnv()

	err := firebase.InitFirebase()
	if err != nil {
		level.Error(logger).Log("msg", "Firebase initialization failed", "err", err)
		os.Exit(1)
	}

	dbConn := config.GetDBConn()
	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		level.Error(logger).Log("msg", "Database connection failed", "err", err)
		os.Exit(1)
	}

	db.AutoMigrate(&model.CaseStudy{})

	repository := casestudy.NewRepository(db)
	service := casestudy.NewService(repository)
	endpoints := casestudy.MakeEndpoints(service)

	httpHandler := transport.NewHTTPHandler(endpoints, logger)

	// CORS ayarlarını yapılandır
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// CORS middleware'ini router'a ekleyin
	corsHandler := handlers.CORS(corsHeaders, corsOrigins, corsMethods)

	// HTTP sunucusunu başlat
	http.Handle("/", corsHandler(httpHandler))
	level.Info(logger).Log("msg", "HTTP server is running", "port", "8080")
	http.ListenAndServe(":8080", nil)
	

}
