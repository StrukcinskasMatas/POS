package database

import (
	"database/sql"
	"fmt"
	"os"

	"pos/utils"

	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

const DB_DRIVER_NAME = "postgres"

type Database struct {
	DB *sql.DB
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s sslmode=disable user=%s password=%s dbname=%s",
		os.Getenv("POS_DB_HOST"), os.Getenv("POS_DB_PORT"), os.Getenv("POS_DB_USER"),
		os.Getenv("POS_DB_PASS"), os.Getenv("POS_DB_NAME"))
}

func CreateDatabaseInstance(connectionString string) *Database {
	database, err := sql.Open(DB_DRIVER_NAME, connectionString)
	if err != nil {
		panic(fmt.Sprintf("Failed creating database instance: %s", err))
	}

	utils.Logger.Infow("Successfully created database instance.")
	return &Database{
		DB: database,
	}
}

func (db *Database) MustPrepare(query string) *Statement {
	statement, err := db.DB.Prepare(query)
	if err != nil {
		panic(fmt.Sprintf("Failed preparing statement: %s", err))
	}

	return &Statement{statement}
}

func (db *Database) CloseConnection() {
	err := db.DB.Close()
	if err != nil {
		utils.Logger.Errorw("Closing database connection failed.")
	}
}

func (db *Database) GetDriver() (database.Driver, error) {
	return postgres.WithInstance(db.DB, &postgres.Config{})
}
