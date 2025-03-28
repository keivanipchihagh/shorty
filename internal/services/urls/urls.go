package urls

import (
	"github.com/keivanipchihagh/shorty/pkg/models"
	"github.com/keivanipchihagh/shorty/pkg/repositories"
)

type UrlService struct {
	UrlRepo repositories.UrlRepo
}

func NewUrlService(UrlRepo repositories.UrlRepo) UrlService {
	return UrlService{UrlRepo: UrlRepo}
}

func (s *UrlService) Create(*models.URL) (*models.URL, error) {
	return nil, nil
}

func (s *UrlService) GetById(id int) (*models.URL, error) {
	url, err := s.UrlRepo.GetById(id)
	return url, err
}

func (s *UrlService) GetAll() ([]models.URL, error) {
	urls, err := s.UrlRepo.GetAll()
	return urls, err
}
