package repositories

import (
	"context"
	"fmt"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/app/models"
	"pmhb-book-service/internal/kerrors"
	"pmhb-book-service/internal/pkg/db/mssqldb"
	"pmhb-book-service/internal/pkg/klog"
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

	//BookRepository groups all function integrate with transaction collection in mssqldbdb
	BookRepository interface {
		GetBook(ctx context.Context, req models.GetBookRepoReq) (models.Book, error)
		//InsertTransaction(ctx context.Context, req models.InsertTransactionRepoReq) (int64, error)
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

// GetBook function
func (tr *BookRepo) GetBook(ctx context.Context, req models.GetBookRepoReq) (models.Book, error) {
	var book models.Book
	ctx = context.Background()

	err := tr.s.Database.PingContext(ctx)
	if err != nil {
		return book, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
	}

	db := fmt.Sprintf("%s.%s.%s", tr.c.MSSQL.DatabaseName, "dbo", tr.c.MSSQL.Tables.Transactions)
	tsql := fmt.Sprintf("SELECT id, name, author FROM %s WHERE id=%s;", db, req.ID)
	fmt.Println("tsql:", tsql)
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

//
//// InsertTransaction function
//func (tr *TransactionsRepo) InsertTransaction(ctx context.Context, req models.InsertTransactionRepoReq) (int64, error) {
//	var newID int64
//	ctx = context.Background()
//
//	err := tr.s.Database.PingContext(ctx)
//	if err != nil {
//		return newID, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
//	}
//	db := fmt.Sprintf("%s.%s.%s", tr.c.MSSQL.DatabaseName, "dbo", tr.c.MSSQL.Tables.Transactions)
//	tsql := fmt.Sprintf("INSERT INTO %s (transaction_name) VALUES (@transaction_name); select convert(bigint, SCOPE_IDENTITY());", db)
//
//	stmt, err := tr.s.Database.Prepare(tsql)
//	if err != nil {
//		return newID, tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
//	}
//	defer stmt.Close()
//
//	row := stmt.QueryRowContext(
//		ctx,
//		sql.Named("transaction_name", req.TransactionName),
//	)
//
//	err = row.Scan(&newID)
//	if err != nil {
//		return newID, tr.errRepo.Wrap(err, kerrors.DatabaseScanErr, nil)
//	}
//
//	return newID, nil
//}

// // UpdateOneTransaction function
// func (tr *TransactionsRepo) UpdateOneTransaction(ctx context.Context, search, update bson.M) error {
// 	count, err := tr.s.Database(tr.c.MSSQL.DatabaseName).Collection(tr.c.MSSQL.Tables.Transactions).UpdateOne(
// 		context.Background(),
// 		search,
// 		update)
// 	if err != nil {
// 		return tr.errRepo.Wrap(err, kerrors.DatabaseServerError, nil)
// 	}
// 	if count.ModifiedCount == 0 {
// 		return tr.errRepo.Wrap(errors.New("No data has been found"), kerrors.NotFoundItemInQuery, nil)
// 	}
// 	return nil
// }
