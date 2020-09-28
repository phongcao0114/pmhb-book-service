package api

import (
	"net/http"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/app/handlers"
	"pmhb-book-service/internal/app/repositories"
	"pmhb-book-service/internal/app/services"

	"pmhb-book-service/internal/pkg/db/mssqldb"
)

type (
	middleware = func(http.Handler) http.Handler
	route      struct {
		desc        string
		path        string
		method      string
		handler     http.HandlerFunc
		middlewares []middleware
	}
)

// CreateBookHandler function
func CreateBookHandler(conf *config.Configs, dbconn *mssqldb.MSSQLConnections) *handlers.BookHandler {
	repo := repositories.NewBookRepo(conf, dbconn)
	srv := services.NewBookService(conf, repo)
	return handlers.NewBookHandler(conf, srv)
}
