package casestudy

import (
	"backend/pkg/casestudy"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	router := mux.NewRouter()

	router.HandleFunc("/casestudy", createCaseStudy(service)).Methods("POST")
	router.HandleFunc("/casestudy/getall", getAllCaseStudy(service)).Methods("GET")
	router.HandleFunc("/casestudy/{id}", getCaseStudy(service)).Methods("GET")

	return router
}

func createCaseStudy(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cs casestudy.CaseStudy
		if err := json.NewDecoder(r.Body).Decode(&cs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.Create(&cs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cs)
	}
}

func getCaseStudy(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseUint(idStr, 10, 32) // 10: decimal, 32: bit size
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		cs, err := s.Get(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getAllCaseStudy(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := s.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
