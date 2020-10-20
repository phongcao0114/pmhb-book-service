package handlers

import (
	"net/http"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/app/response"
	"pmhb-book-service/internal/app/services"
	"pmhb-book-service/internal/app/utils"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/klog"
	"pmhb-book-service/models"

	"github.com/gorilla/mux"
)

const (
	// BookHandlerPrefix prefix logger
	BookHandlerPrefix = "Book_handler"
)

// BookHandler struct defines the variables for specifying interface.
type BookHandler struct {
	conf       *config.Configs
	errHandler kerrors.KError
	logger     klog.Logger

	srv services.BookService
}

// NewBookHandler connects to the service from handler.
func NewBookHandler(conf *config.Configs, s services.BookService) *BookHandler {
	return &BookHandler{
		conf:       conf,
		errHandler: kerrors.WithPrefix(BookHandlerPrefix),
		logger:     klog.WithPrefix(BookHandlerPrefix),

		srv: s,
	}
}

// GetBook handler handles the upcoming request.
func (th *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	commitModels, err := th.srv.GetBookByID(r.Context(), id)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	response.WriteJSON(w)(response.HandleSuccess(r, commitModels))
	return

}

// GetBook handler handles the upcoming request.
func (th *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	commitModels, err := th.srv.GetBooks(r.Context())
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	response.WriteJSON(w)(response.HandleSuccess(r, commitModels))
	return

}

// InsertBook handler handles the upcoming request.
func (th *BookHandler) InsertBook(w http.ResponseWriter, r *http.Request) {
	var body models.InsertBookReq
	err := utils.DecodeToBody(&th.errHandler, &body, r)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}
	id, err := th.srv.InsertBook(r.Context(), models.InsertBookReq{
		Name:   body.Name,
		Author: body.Author,
	})
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	response.WriteJSON(w)(response.HandleSuccess(r, id))
	return
}

func (th *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var body models.InsertBookReq
	err := utils.DecodeToBody(&th.errHandler, &body, r)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	err = th.srv.UpdateBook(r.Context(), id, models.UpdateBookReq{
		Name:   body.Name,
		Author: body.Author,
	})
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	response.WriteJSON(w)(response.HandleSuccess(r, true))
	return
}

func (th *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := th.srv.DeleteBook(r.Context(), id)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	response.WriteJSON(w)(response.HandleSuccess(r, true))
	return
}
