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
		GetBookByID(ctx context.Context, id string) (models.Book, error)
		InsertBook(ctx context.Context, req models.InsertBookReq) (string, error)
		GetBooks(ctx context.Context) ([]models.Book, error)
		UpdateBook(ctx context.Context, id string, req models.UpdateBookReq) error
		DeleteBook(ctx context.Context, id string) error
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
func (tr *BookSrv) GetBookByID(ctx context.Context, id string) (models.Book, error) {
	return tr.bookRepo.GetBookByID(ctx, id)
}

// InsertBook func
func (tr *BookSrv) InsertBook(ctx context.Context, req models.InsertBookReq) (string, error) {
	return tr.bookRepo.InsertBook(ctx, req)
}

// InsertBook func
func (tr *BookSrv) GetBooks(ctx context.Context) ([]models.Book, error) {
	return tr.bookRepo.GetBooks(ctx)
}

func (tr *BookSrv) UpdateBook(ctx context.Context, id string, req models.UpdateBookReq) error {
	return tr.bookRepo.UpdateBook(ctx, id, req)
}

func (tr *BookSrv) DeleteBook(ctx context.Context, id string) error {
	return tr.bookRepo.DeleteBook(ctx, id)
}
