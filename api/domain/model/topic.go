package model

import "time"

type Topic struct {
	TopicID     int        `db:"topic_id"`
	UserID      int        `db:"user_id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
