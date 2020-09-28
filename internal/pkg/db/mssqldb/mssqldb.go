package mssqldb

import (
	"context"
	"database/sql"
	"fmt"
	"pmhb-book-service/internal/app/config"
	"pmhb-book-service/internal/pkg/klog"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

type (
	// MSSQLConnections hold all MSSQLDB connections that needed for the app
	MSSQLConnections struct {
		Database *sql.DB
	}
)

// NewDatabaseConnection connects to MSSQLDB
func NewDatabaseConnection(conf *config.Configs) (*MSSQLConnections, error) {
	db, err := DialInfo(conf)
	if err != nil {
		klog.WithPrefix("mssqldb").Errorf("can't connect to database err: %v", err)
		return nil, err
	}

	msSQLConn := &MSSQLConnections{
		Database: db,
	}

	return msSQLConn, nil
}

// DialInfo dial to the target server
func DialInfo(clientConf *config.Configs) (*sql.DB, error) {

	KLogger := klog.WithPrefix("MSSQLDB")
	KLogger.Infof("[DialInfo] Dialing to mssql")

	if len(clientConf.MSSQL.MSSQLAddress) == 0 {
		KLogger.Panic("No DB address")
	}
	address := strings.Split(clientConf.MSSQL.MSSQLAddress[0], ":")
	server := address[0]
	port, err := strconv.Atoi(address[1])
	if err != nil {
		KLogger.Panic("Error converting port to number")
	}

	// Create connection string
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, clientConf.MSSQL.Username, clientConf.MSSQL.Password, port)
	// Create connection pool
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		KLogger.Panic("Error creating connection pool: " + err.Error())
	}
	KLogger.Infof("Connected to %v!\n", connectionString)

	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	if err := db.PingContext(ctx); err != nil {
		KLogger.Panic("Error pinging database: " + err.Error())
	}

	// Log output
	KLogger.WithFields(map[string]interface{}{
		"service":   "mssql",
		"addresses": clientConf.MSSQL.MSSQLAddress,
		"database":  clientConf.MSSQL.DatabaseName,
	}).Infoln("Successfully dialing to MSSQL")

	return db, nil
}
