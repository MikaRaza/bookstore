// database/database.go
package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importer le pilote MySQL
)

var DB *sql.DB

func ConnectDatabase() error {
	dsn := os.Getenv("DSN") // Assurez-vous que la variable d'environnement est correctement configurée
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
