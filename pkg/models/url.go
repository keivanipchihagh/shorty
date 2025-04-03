package models

import "time"

/*
	CREATE TABLE urls(
		id BIGINT NOT NULL,
		original text NOT NULL,
		shortened text NOT NULL,
		created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
		expires_at timestamp with time zone,
		PRIMARY KEY(id)
	);
*/

type URL struct {
	ID        int64     `json:"id"`         // Unique identifier for the URL entry
	Original  string    `json:"original"`   // Original URL that is being shortened
	Shortened string    `json:"shortened"`  // Identifier for the shortened URL
	CreatedAt time.Time `json:"created_at"` // Time when the URL entry was created
	ExpiresAt time.Time `json:"expires_at"` // Time when the URL entry is set to expire
}
