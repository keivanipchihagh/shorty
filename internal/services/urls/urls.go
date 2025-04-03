package urls

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/keivanipchihagh/shorty/internal/services/kgs"
	"github.com/keivanipchihagh/shorty/pkg/models"
	"github.com/keivanipchihagh/shorty/pkg/repositories"
)

type UrlService struct {
	UrlRepo     repositories.UrlRepo
	RedisClient *redis.Client
}

func NewUrlService(urlRepo repositories.UrlRepo, redisClient *redis.Client) UrlService {
	return UrlService{
		UrlRepo:     urlRepo,
		RedisClient: redisClient,
	}
}

func (s *UrlService) Create(url *models.URL) error {

	id, shortened, err := kgs.GenerateId()
	if err != nil {
		return err
	}
	url.ID = id
	url.Shortened = shortened

	url.CreatedAt = time.Now()
	url.ExpiresAt = time.Now().Add(time.Hour)

	err = s.UrlRepo.Create(url)
	return err
}

func (s *UrlService) GetById(id int64) (*models.URL, error) {
	url, err := s.UrlRepo.GetById(id)
	return url, err
}

func (s *UrlService) GetByShortened(shortened string) (*models.URL, error) {

	// Check if there is a value in redis
	result, err := s.RedisClient.Get(context.Background(), shortened).Result()
	if err == nil {
		url := &models.URL{
			Shortened: shortened,
			Original:  result,
		}
		return url, nil
	}

	// Fetch the original URL from database
	url, err := s.UrlRepo.GetByShortened(shortened)
	if err != nil {
		return nil, err
	}

	// Save the value in redis
	err = s.RedisClient.Set(context.Background(), shortened, url.Original, time.Minute).Err()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *UrlService) GetAll() ([]models.URL, error) {
	urls, err := s.UrlRepo.GetAll()
	return urls, err
}
