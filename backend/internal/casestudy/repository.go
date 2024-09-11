package casestudy

import (
	"backend/pkg/casestudy"

	"gorm.io/gorm"
)

type Repository interface {
	Create(cs *casestudy.CaseStudy) error
	Get(id uint) (*casestudy.CaseStudy, error)
	GetAll() ([]*casestudy.CaseStudy, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(cs *casestudy.CaseStudy) error {
	return r.db.Create(cs).Error
}

func (r *repository) Get(id uint) (*casestudy.CaseStudy, error) {
	var cs casestudy.CaseStudy
	if err := r.db.First(&cs, id).Error; err != nil {
		return nil, err
	}
	return &cs, nil
}

func (r *repository) GetAll() ([]*casestudy.CaseStudy, error) {
	var cs []*casestudy.CaseStudy // Slice of pointers to CaseStudy

	// Veritabanından tüm CaseStudy'leri çekiyoruz
	if err := r.db.Find(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}
