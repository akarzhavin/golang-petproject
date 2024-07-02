package models

import (
	"context"
	"time"
)

// RefreshToken is the structure which holds one refresh token from the database.
type RefreshToken struct {
	Token     string    `json:"token"`
	UserID    int       `json:"user_id"`
	Active    bool      `json:"active"`
	UsedCount int8      `json:"used_count"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (token *RefreshToken) Store() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into auth_refresh_tokens (token, user_id, active, used_count, created_at, expires_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := db.QueryRowContext(ctx, stmt,
		token.Token,
		token.UserID,
		token.Active,
		token.UsedCount,
		token.CreatedAt,
		token.ExpiresAt,
	).Scan(&newID)

	if err != nil {
		return err
	}

	return nil
}

func (token *RefreshToken) GetOne(tokenStr string) (*RefreshToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select token, user_id, active, used_count, created_at, expires_at from auth_refresh_tokens where token = $1`

	var foundToken RefreshToken
	row := db.QueryRowContext(ctx, query, tokenStr)

	err := row.Scan(
		&foundToken.Token,
		&foundToken.UserID,
		&foundToken.Active,
		&foundToken.UsedCount,
		&foundToken.CreatedAt,
		&foundToken.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &foundToken, nil
}
