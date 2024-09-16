package casestudy

import (
	"backend/pkg/casestudy"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, cs *casestudy.CaseStudy) error
	Get(ctx context.Context, id uint) (*casestudy.CaseStudy, error)
	GetAll(ctx context.Context) ([]*casestudy.CaseStudy, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, cs *casestudy.CaseStudy) error {
	return r.db.WithContext(ctx).Create(cs).Error
}

func (r *repository) Get(ctx context.Context, id uint) (*casestudy.CaseStudy, error) {
	var cs casestudy.CaseStudy
	if err := r.db.WithContext(ctx).First(&cs, id).Error; err != nil {
		return nil, err
	}
	return &cs, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*casestudy.CaseStudy, error) {
	var cs []*casestudy.CaseStudy
	if err := r.db.WithContext(ctx).Find(&cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}
