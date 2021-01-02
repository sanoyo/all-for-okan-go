package repository

import (
	"github.com/sanoyo/all-for-okan-go/api/domain/model"
)

type TopicsRepository interface {
	FetchTopics() (*[]model.Topic, error)
	FetchTopic(topicID uint) (*model.Topic, error)
}
