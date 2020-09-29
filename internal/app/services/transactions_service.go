package services

import (
	"context"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/app/models"
	"pmhb-book-service/internal/app/repositories"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/klog"
)

const (
	// BookServicePrefix prefix logger
	BookServicePrefix = "Book_service"
)

type (
	// BookSrv groups all transactions service together
	BookSrv struct {
		conf     *config.Configs
		errSrv   kerrors.KError
		logger   klog.Logger
		bookRepo repositories.BookRepository
	}

	//BookService interface
	BookService interface {
		GetBook(ctx context.Context, req *models.GetBookSrvReq) (models.Book, error)
		//InsertTransaction(ctx context.Context, req *models.InsertTransactionSrvReq) (models.InsertTransactionSrvRes, error)
	}
)

//NewBookService init a new transactions service
func NewBookService(conf *config.Configs, repo repositories.BookRepository) *BookSrv {
	return &BookSrv{
		conf:     conf,
		errSrv:   kerrors.WithPrefix(BookServicePrefix),
		logger:   klog.WithPrefix(BookServicePrefix),
		bookRepo: repo,
	}
}

// BookSrv function service
func (tr *BookSrv) GetBook(ctx context.Context, req *models.GetBookSrvReq) (models.Book, error) {
	return tr.bookRepo.GetBook(ctx, models.GetBookRepoReq{
		ID: req.ID,
	})
}

//
//// InsertTransaction function service
//func (tr *BookSrv) InsertTransaction(ctx context.Context, req *models.InsertTransactionSrvReq) (models.InsertTransactionSrvRes, error) {
//	transID, err := tr.transactionsRepo.InsertTransaction(ctx, models.InsertTransactionRepoReq{
//		TransactionName: req.TransactionName,
//	})
//	if err != nil {
//		return models.InsertTransactionSrvRes{}, err
//	}
//	reqModel := models.InsertTransactionSrvRes{
//		TransactionID:   transID,
//		TransactionName: req.TransactionName,
//	}
//	return reqModel, nil
//}
