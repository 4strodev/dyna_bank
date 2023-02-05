package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/4strodev/dyna-bank/internal/app/shared/errors/utils"
	_ "github.com/lib/pq"
)

var DBPool *sql.DB

func init() {
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	log.Println(uri)
	DBPool, err = sql.Open("postgres", uri)
	utils.FailOnError(err, "Error connecting to postgres")
}
