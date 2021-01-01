package model

import "time"

type Topic struct {
	ID          int        `db:"id"`
	UserID      int        `db:"user_id"`
	Title       int16      `db:"title"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
