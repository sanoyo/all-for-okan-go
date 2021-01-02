package usecase

import (
	"github.com/sanoyo/all-for-okan-go/api/domain/model"
	"github.com/sanoyo/all-for-okan-go/api/domain/repository"
)

type TopicsUseCase interface {
	FetchTopics() (*[]model.Topic, error)
	FetchTopic(topicID uint) (*model.Topic, error)
}

type topicsUseCaseImpl struct {
	topicRepo repository.TopicsRepository
}

func NewTopicsUseCase(topicRepo repository.TopicsRepository) TopicsUseCase {
	return &topicsUseCaseImpl{
		topicRepo: topicRepo,
	}
}

func (a *topicsUseCaseImpl) FetchTopics() (*[]model.Topic, error) {
	topics, err := a.topicRepo.FetchTopics()
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (a *topicsUseCaseImpl) FetchTopic(topicID uint) (*model.Topic, error) {
	topic, err := a.topicRepo.FetchTopic(topicID)
	if err != nil {
		return nil, err
	}

	return topic, nil
}
