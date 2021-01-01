package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoyo/all-for-okan-go/api/usecase"
)

type TopicsHandler interface {
	Get(*gin.Context)
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
