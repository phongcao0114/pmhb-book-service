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
)

const (
	// TransactionHandlerPrefix prefix logger
	TransactionHandlerPrefix = "Transaction_handler"
)

// TransactionHandler struct defines the variables for specifying interface.
type BookHandler struct {
	conf       *config.Configs
	errHandler kerrors.KError
	logger     klog.Logger

	srv services.BookService
}

// NewTransactionHandler connects to the service from handler.
func NewBookHandler(conf *config.Configs, s services.BookService) *BookHandler {
	return &BookHandler{
		conf:       conf,
		errHandler: kerrors.WithPrefix(TransactionHandlerPrefix),
		logger:     klog.WithPrefix(TransactionHandlerPrefix),

		srv: s,
	}
}

// GetBook handler handles the upcoming request.
func (th *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {

	var body models.GetBookReq
	err := utils.DecodeToBody(&th.errHandler, &body, r)
	if err != nil {
		response.WriteJSON(w)(response.HandleError(r, err))
		return
	}

	commitModels, err := th.srv.GetBook(r.Context(), &models.GetBookSrvReq{
		ID: body.ID,
	})
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

	id, err := th.srv.InsertBook(r.Context(), &models.InsertBookReq{
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
