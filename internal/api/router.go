package api

import (
	"context"
	"encoding/json"
	"net/http"
	"pmhb-book-service/internal/app/config"

	"pmhb-book-service/internal/pkg/db/mssqldb"
	"pmhb-book-service/internal/pkg/middlewares"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

// NewRouter return new mux router with a closer for cleaning up underlying resources
func NewRouter(conf *config.Configs, dbconn *mssqldb.MSSQLConnections) (*mux.Router, error) {

	// Book handler API
	bookHandler := CreateBookHandler(conf, dbconn)

	router := mux.NewRouter()

	// the place to define all route we need
	r := []route{
		{
			desc:   "API for checking connection",
			method: get,
			path:   "/ping",
			handler: func(w http.ResponseWriter, r *http.Request) {
				JSON(r.Context(), w, http.StatusOK, map[string]interface{}{"data": "pong"})
				return
			},
		},
		{
			desc:    "Get API for payment hub",
			method:  get,
			path:    "/kph/api/book/get",
			handler: bookHandler.GetBook,
		},
		//{
		//	desc:    "Set API for payment hub",
		//	method:  post,
		//	path:    "/kph/api/set",
		//	handler: transactionHandler.InsertTransaction,
		//},
	}
	router.Use(middlewares.Recover)
	router.Use(middlewares.AcceptLanguage)
	router.Use(middlewares.LoggerWithRequestMeta)
	router.Use(middlewares.RequestInfo)
	// the for loop to add router in to mux router
	for _, rte := range r {
		router.Path(rte.path).Methods(rte.method).HandlerFunc(rte.handler)
	}
	return router, nil
}

// AppError interface
type AppError interface {
	GetHTTPStatus() int
	WithContext(ctx context.Context) error
	Error() string
}

// JSON response
func JSON(ctx context.Context, w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		Error(ctx, w, errors.Wrap(err, "JSON marshal failed"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(b)
}

// Error main function
func Error(ctx context.Context, w http.ResponseWriter, err error, status int) {
	if appErr, ok := err.(AppError); ok {
		JSON(ctx, w, appErr.GetHTTPStatus(), appErr.WithContext(ctx))
		return
	}
	JSON(ctx, w, status, errors.New("internal server error"))
}
