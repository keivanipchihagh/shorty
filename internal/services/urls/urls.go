package urls

import (
	"time"

	"github.com/keivanipchihagh/shorty/pkg/models"
	"github.com/keivanipchihagh/shorty/pkg/repositories"
)

type UrlService struct {
	UrlRepo repositories.UrlRepo
}

func NewUrlService(UrlRepo repositories.UrlRepo) UrlService {
	return UrlService{UrlRepo: UrlRepo}
}

func (s *UrlService) Create(url *models.URL) error {

	url.CreatedAt = time.Now()
	url.ExpiresAt = time.Now().Add(time.Hour)

	err := s.UrlRepo.Create(url)
	return err
}

func (s *UrlService) GetById(id int) (*models.URL, error) {
	url, err := s.UrlRepo.GetById(id)
	return url, err
}

func (s *UrlService) GetAll() ([]models.URL, error) {
	urls, err := s.UrlRepo.GetAll()
	return urls, err
}
