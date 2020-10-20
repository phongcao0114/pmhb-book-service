package repositories

import (
	"context"
	"errors"
	"fmt"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/db/mariadb"
	"pmhb-book-service/internal/pkg/klog"
	"pmhb-book-service/models"

	"github.com/google/uuid"
)

const (
	// BookRepositoryPrefix prefix repo
	BookRepositoryPrefix = "Book_repository"
)

type (
	// BookRepo defines mariadb.Client connection for each client
	BookRepo struct {
		s       *mariadb.MariaDBConnections
		c       *config.Configs
		errRepo kerrors.KError
		logger  klog.Logger
	}

	//BookRepository groups all function integrate with book collection in mariadb
	BookRepository interface {
		GetBookByID(ctx context.Context, id string) (models.Book, error)
		GetBooks(ctx context.Context) ([]models.Book, error)
		InsertBook(ctx context.Context, req models.InsertBookReq) (string, error)
		UpdateBook(ctx context.Context, id string, req models.UpdateBookReq) error
		DeleteBook(ctx context.Context, id string) error
	}
)

// NewBookRepo opens the connection to DB from repositories package
func NewBookRepo(configs *config.Configs, s *mariadb.MariaDBConnections) *BookRepo {
	// Return model
	return &BookRepo{
		s:       s,
		c:       configs,
		errRepo: kerrors.WithPrefix(BookRepositoryPrefix),
		logger:  klog.WithPrefix(BookRepositoryPrefix),
	}
}

// GetBookByID function
func (tr *BookRepo) GetBookByID(ctx context.Context, id string) (models.Book, error) {
	var book models.Book
	ctx = context.Background()
	stmt, err := tr.s.Database.Prepare("SELECT * FROM book WHERE id=?")
	if err != nil {
		return book, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return book, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
	}
	for rows.Next() {
		if err = rows.Scan(&book.ID, &book.Name, &book.Author); err != nil {
			return book, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
		}
	}
	return book, nil
}

//GetBooks function
func (tr *BookRepo) GetBooks(ctx context.Context) ([]models.Book, error) {
	var bookList []models.Book
	stmt, err := tr.s.Database.Prepare("SELECT * FROM book")
	if err != nil {
		return bookList, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	ctx = context.Background()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return bookList, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
	}
	defer stmt.Close()
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.ID, &book.Name, &book.Author); err != nil {
			return bookList, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
		}
		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (tr *BookRepo) InsertBook(ctx context.Context, req models.InsertBookReq) (string, error) {
	if req.Name == "" {
		return "", tr.errRepo.Wrap(errors.New("name is missing"), kerrors.ValidateFailed, nil)
	}
	if req.Author == "" {
		return "", tr.errRepo.Wrap(errors.New("author is missing"), kerrors.ValidateFailed, nil)
	}
	id := uuid.New().String()
	ctx = context.Background()

	stmt, err := tr.s.Database.Prepare("INSERT INTO book VALUES(?,?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id, req.Name, req.Author)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (tr *BookRepo) UpdateBook(ctx context.Context, id string, req models.UpdateBookReq) error {
	tsql := fmt.Sprintf("UPDATE book SET ")
	if req.Name != "" {
		tsql = tsql + fmt.Sprintf("name='%s',", req.Name)
	}
	if req.Author != "" {
		tsql = tsql + fmt.Sprintf("author='%s',", req.Author)
	}
	tsql = tsql[:len(tsql)-1]
	tsql = tsql + fmt.Sprintf(" WHERE id='%s'", id)
	fmt.Println("tsql:", tsql)
	stmt, err := tr.s.Database.Prepare(tsql)
	if err != nil {
		return tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	defer stmt.Close()
	ctx = context.Background()
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
	}
	return nil
}

func (tr *BookRepo) DeleteBook(ctx context.Context, id string) error {
	stmt, err := tr.s.Database.Prepare("DELETE FROM book WHERE id=?")
	if err != nil {
		return tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	defer stmt.Close()
	ctx = context.Background()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	return nil
}
