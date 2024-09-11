package casestudy

import (
	"backend/pkg/casestudy"
)

type Service interface {
	Create(cs *casestudy.CaseStudy) error
	Get(id uint) (*casestudy.CaseStudy, error)
	GetAll() ([]*casestudy.CaseStudy, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(cs *casestudy.CaseStudy) error {
	return s.repo.Create(cs)
}

func (s *service) Get(id uint) (*casestudy.CaseStudy, error) {
	return s.repo.Get(id)
}

func (s *service) GetAll() ([]*casestudy.CaseStudy, error) {
	return s.repo.GetAll()
}
