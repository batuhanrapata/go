package casestudy

import (
	"backend/pkg/casestudy"
	"context"
)

type Service interface {
	Create(ctx context.Context, cs *casestudy.CaseStudy) error
	Get(ctx context.Context, id uint) (*casestudy.CaseStudy, error)
	GetAll(ctx context.Context) ([]*casestudy.CaseStudy, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, cs *casestudy.CaseStudy) error {
	return s.repo.Create(ctx, cs)
}

func (s *service) Get(ctx context.Context, id uint) (*casestudy.CaseStudy, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]*casestudy.CaseStudy, error) {
	return s.repo.GetAll(ctx)
}
