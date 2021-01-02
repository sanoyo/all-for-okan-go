package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanoyo/all-for-okan-go/api/usecase"
)

type TopicsHandler interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
}

type topicsHandler struct {
	useCase usecase.TopicsUseCase
}

func NewTopicsHandler(useCase usecase.TopicsUseCase) TopicsHandler {
	return &topicsHandler{
		useCase: useCase,
	}
}

func (a *topicsHandler) Get(c *gin.Context) {
	topics, err := a.useCase.FetchTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (a *topicsHandler) GetByID(c *gin.Context) {
	topicIDInt, err := strconv.Atoi(c.Param("topicID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	topicID := uint(topicIDInt)

	topic, err := a.useCase.FetchTopic(topicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topic)
}
