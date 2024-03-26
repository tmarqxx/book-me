package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func OpenEmbeddedDB(dbFile string, dbUrl string, authToken string) (*sql.DB, string) {
	dir, err := os.MkdirTemp("", "libsql-*")

	if err != nil {
		log.Fatal("Error creating temporary directory: ", err)
	}

	dbPath := filepath.Join(dir, dbFile)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbUrl, libsql.WithAuthToken(authToken))
	if err != nil {
		os.RemoveAll(dir)
		log.Fatal("Error creating connector: ", err)
	}

	db := sql.OpenDB(connector)
	e := db.Ping()
	if e != nil {
		db.Close()
		log.Fatal("Error while trying to ping DB: ", err)
	}

	return db, dir
}
