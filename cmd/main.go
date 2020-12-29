package main

import (
	"net/http"
	"os"

	"github.com/sanoyo/all-for-okan-go/api/infra/persistence"
)

const (
	port    = "8080"
	maxconn = 10
)

func main() {
	db.Init(maxconn)

	var dbx *sqlx.DB
	dbx, err := sqlx.Connect("pgx", config.DB.DSN)
	if err != nil {
		log.Warn("failed to setup database", "err", err)
		return
	}

	topicPersistence := persistence.NewTopicsPersistence()
	topicUseCase := usecase.NewTopicUseCase(topicPersistence)
	topicHandler := handler.NewTopicHandler(topicUseCase)

	router := newRouter()

	// BHI保守ユーザー用エンドポイント
	router.Route("/topic", func(r chi.Router) {
		r.Get("/topic/{topicId}", topicHandler.Get)
		r.Post("/topic/create", topicHandler.Create)
	})

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.ErrorWithReport(err, "Listen and serve failed.", "err", err)
		os.Exit(1)
	}
}

func newRouter() *chi.Mux {
	// middlewareの設定
	router := chi.NewRouter()
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	return router
}
