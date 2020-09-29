package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/db/mssqldb"
	"pmhb-book-service/internal/pkg/klog"
	"pmhb-book-service/models"

	"github.com/google/uuid"
)

const (
	// BookRepositoryPrefix prefix repo
	BookRepositoryPrefix = "Book_repository"
)

type (
	// BookRepo defines mssqldb.Client connection for each client
	BookRepo struct {
		s       *mssqldb.MSSQLConnections
		c       *config.Configs
		errRepo kerrors.KError
		logger  klog.Logger
	}

	//BookRepository groups all function integrate with book collection in mssqldbdb
	BookRepository interface {
		GetBookByID(ctx context.Context, req models.GetBookRepoReq) (models.Book, error)
		InsertBook(ctx context.Context, req models.InsertBookReq) (string, error)
	}
)

// NewBookRepo opens the connection to DB from repositories package
func NewBookRepo(configs *config.Configs, s *mssqldb.MSSQLConnections) *BookRepo {
	// Return model
	return &BookRepo{
		s:       s,
		c:       configs,
		errRepo: kerrors.WithPrefix(BookRepositoryPrefix),
		logger:  klog.WithPrefix(BookRepositoryPrefix),
	}
}

// GetBookByID function
func (tr *BookRepo) GetBookByID(ctx context.Context, req models.GetBookRepoReq) (models.Book, error) {
	var book models.Book
	ctx = context.Background()

	err := tr.s.Database.PingContext(ctx)
	if err != nil {
		return book, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}

	db := fmt.Sprintf("%s.%s.%s", tr.c.MSSQL.DatabaseName, "dbo", tr.c.MSSQL.Tables.Transactions)
	tsql := fmt.Sprintf("SELECT id, name, author FROM %s WHERE id='%s';", db, req.ID)
	// Execute query
	rows, err := tr.s.Database.QueryContext(ctx, tsql)
	if err != nil {
		return book, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
	}

	defer rows.Close()
	// Iterate through the result set.
	for rows.Next() {
		// Get values from row.
		err := rows.Scan(&book.ID, &book.Name, &book.Author)
		if err != nil {
			return book, tr.errRepo.Wrap(err, kerrors.CannotGetDataFromDB, nil)
		}
		//dbList = append(dbList, books)
	}

	return book, nil
}

func (tr *BookRepo) InsertBook(ctx context.Context, req models.InsertBookReq) (string, error) {
	id := uuid.New().String()
	ctx = context.Background()

	err := tr.s.Database.PingContext(ctx)
	if err != nil {
		return "", tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	db := fmt.Sprintf("%s.%s.%s", tr.c.MSSQL.DatabaseName, "dbo", tr.c.MSSQL.Tables.Transactions)
	tsql := fmt.Sprintf("INSERT INTO %s VALUES (@id,@name,@author)", db)
	stmt, err := tr.s.Database.Prepare(tsql)
	if err != nil {
		return "", tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		sql.Named("id", id),
		sql.Named("name", req.Name),
		sql.Named("author", req.Author),
	)
	if err != nil {
		return "", tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}
	return id, nil
}
