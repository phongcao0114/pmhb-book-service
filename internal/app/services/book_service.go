package services

import (
	"context"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/app/repositories"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/klog"
	"pmhb-book-service/models"
)

const (
	// BookServicePrefix prefix logger
	BookServicePrefix = "Book_service"
)

type (
	// BookSrv groups all book service together
	BookSrv struct {
		conf     *config.Configs
		errSrv   kerrors.KError
		logger   klog.Logger
		bookRepo repositories.BookRepository
	}

	//BookService interface
	BookService interface {
		GetBookByID(ctx context.Context, req models.GetBookSrvReq) (models.Book, error)
		InsertBook(ctx context.Context, req models.InsertBookReq) (string, error)
	}
)

//NewBookService init a new book service
func NewBookService(conf *config.Configs, repo repositories.BookRepository) *BookSrv {
	return &BookSrv{
		conf:     conf,
		errSrv:   kerrors.WithPrefix(BookServicePrefix),
		logger:   klog.WithPrefix(BookServicePrefix),
		bookRepo: repo,
	}
}

// GetBook func
func (tr *BookSrv) GetBookByID(ctx context.Context, req models.GetBookSrvReq) (models.Book, error) {
	return tr.bookRepo.GetBookByID(ctx, models.GetBookRepoReq{
		ID: req.ID,
	})
}

func (tr *BookSrv) InsertBook(ctx context.Context, req models.InsertBookReq) (string, error) {
	return tr.bookRepo.InsertBook(ctx, models.InsertBookReq{
		Name:   req.Name,
		Author: req.Author,
	})
}
