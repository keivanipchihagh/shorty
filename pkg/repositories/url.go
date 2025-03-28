package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/keivanipchihagh/shorty/pkg/models"
)

type UrlRepo interface {
	Create(url *models.URL) error
	GetById(id int) (*models.URL, error)
	GetAll() ([]models.URL, error)
}

type UrlRepoImp struct {
	db *pgxpool.Pool
}

func NewUrlRepo(db *pgxpool.Pool) UrlRepo {
	return &UrlRepoImp{db: db}
}

func (r *UrlRepoImp) Create(url *models.URL) error {
	query := `
		INSERT INTO urls (original, shortened, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(context.Background(), query, url.Original, url.Shortened, url.ExpiresAt)
	return err
}

func (r *UrlRepoImp) GetById(id int) (*models.URL, error) {
	query := `
		SELECT id, original, shortened, created_at, expires_at
		FROM urls
		WHERE id = $1
	`
	row := r.db.QueryRow(context.Background(), query, id)

	var url models.URL
	if err := row.Scan(
		&url.ID,
		&url.Original,
		&url.Shortened,
		&url.CreatedAt,
		&url.ExpiresAt,
	); err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *UrlRepoImp) GetAll() ([]models.URL, error) {
	query := `
		SELECT id, original, shortened, created_at, expires_at
		FROM urls
	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving URLs: %v", err)
	}
	defer rows.Close()

	var urls []models.URL
	for rows.Next() {
		var url models.URL
		if err := rows.Scan(
			&url.ID,
			&url.Original,
			&url.Shortened,
			&url.CreatedAt,
			&url.ExpiresAt,
		); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
