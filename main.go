package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanoyo/all-for-okan-go/api/handler"
	infrastructure "github.com/sanoyo/all-for-okan-go/api/infra"
	"github.com/sanoyo/all-for-okan-go/api/infra/persistence"
	"github.com/sanoyo/all-for-okan-go/api/usecase"
)

func main() {
	topicPersistence := persistence.NewTopicsPersistence(infrastructure.DB)
	topicUseCase := usecase.NewTopicsUseCase(topicPersistence)
	topicHandler := handler.NewTopicsHandler(topicUseCase)

	r := gin.Default()
	// TODO: CORS実装
	// r.Use(middleware.CORSHeaders())

	r.GET("/topics", func(c *gin.Context) { topicHandler.Get(c) })

	if err := r.Run(); err != nil {
		log.Fatalf("main error: %s", err.Error())
	}
}
