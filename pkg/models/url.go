package models

import "time"

type URL struct {
	ID        int64     `json:"id"`         // Unique identifier for the URL entry
	Original  string    `json:"original"`   // Original URL that is being shortened
	Shortened string    `json:"shortened"`  // Identifier for the shortened URL
	CreatedAt time.Time `json:"created_at"` // Time when the URL entry was created
	ExpiresAt time.Time `json:"expires_at"` // Time when the URL entry is set to expire
}
