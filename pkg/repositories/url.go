package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/keivanipchihagh/shorty/pkg/models"
)

type RepositoryInterface interface {
	Create(url *models.URL) (*models.URL, error)
	GetById(id int) (*models.URL, error)
	GetAll() ([]models.URL, error)
}

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(url *models.URL) (*models.URL, error) {
	query := `
		INSERT INTO urls (original, shortened)
		VALUES ($1, $2)
		RETURNING id, original, shortened, created_at, expires_at
	`
	row := r.db.QueryRow(context.Background(), query, url.Original, url.Shortened)

	if err := row.Scan(&url.ID, &url.Original, &url.Shortened, &url.CreatedAt, &url.ExpiresAt); err != nil {
		return nil, fmt.Errorf("error creating URL: %v", err)
	}

	return url, nil
}

func (r *Repository) GetById(id int) (*models.URL, error) {
	query := `
		SELECT id, original, shortened, created_at, expires_at
		FROM urls
		WHERE id = $1
	`
	row := r.db.QueryRow(context.Background(), query, id)

	var url models.URL
	if err := row.Scan(&url.ID, &url.Original, &url.Shortened, &url.CreatedAt, &url.ExpiresAt); err != nil {
		return nil, fmt.Errorf("error reading URL")
	}

	return &url, nil
}

func (r *Repository) GetAll() ([]models.URL, error) {
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
		if err := rows.Scan(&url.ID, &url.Original, &url.Shortened, &url.CreatedAt, &url.ExpiresAt); err != nil {
			return nil, fmt.Errorf("error scanning URL: %v", err)
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return urls, nil
}
