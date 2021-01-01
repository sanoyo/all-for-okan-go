package persistence

import (
	"github.com/jmoiron/sqlx"
	"github.com/sanoyo/all-for-okan-go/api/domain/model"
)

type topicsPersistence struct {
	DB *sqlx.DB
}

func NewTopicsPersistence(db *sqlx.DB) *topicsPersistence {
	return &topicsPersistence{DB: db}
}

func (app *topicsPersistence) FetchTopics() (*[]model.Topic, error) {
	const sql = `
    SELECT
        id
      , title
	  , description
	  , created_at
	  , updated_at
    FROM
        topics
    `

	var topic []model.Topic
	err := app.DB.Select(&topic, sql)
	if err != nil {
		return nil, err
	}

	return &topic, nil
}
