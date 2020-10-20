package mariadb

import (
	"context"
	"database/sql"
	"log"
	"pmhb-book-service/internal/app/config"

	_ "github.com/go-sql-driver/mysql"
)

type (
	// MariaDBConnections hold all MariaDB connections that needed for the app
	MariaDBConnections struct {
		Database *sql.DB
	}
)

// NewDatabaseConnection connects to MariaDB
func NewDatabaseConnection(conf *config.Configs) (*MariaDBConnections, error) {
	db, err := sql.Open("mysql", "root:12345678x@X@/book")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	mariaConn := &MariaDBConnections{
		Database: db,
	}

	return mariaConn, nil
}
