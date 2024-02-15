package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// DatabaseURL represents the connection string to the database
const DatabaseURL = "host=db user=postgres port=8001 dbname=postgres password=test sslmode=disable"

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	// connStr := "host=db user=postgres port=8001 dbname=postgres password=test sslmode=disable"
	once.Do(func() {
		// Initialize the database connection
		conn, err := sql.Open("postgres", DatabaseURL)
		if err != nil {
			log.Fatal(err)
		}
		db = conn

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to the database")
	})
	return db
	// connStr := "user=postgres dbname=postgres password=test sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := db.Ping(); err != nil {
	// 	return nil, err
	// }

	// return &PostgresStore{
	// 	db: db,
	// }, nil
}

func Init() error {
	newdataBase := GetDB()
	if err := createAccountTable(newdataBase); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createAccountTable(s *sql.DB) error {
	query := `create table if not exists account (
		id serial primary key,
        email varchar(100),
		first_name varchar(100),
		last_name varchar(100),
		number serial,
		encrypted_password varchar(100),
		balance serial,
		created_at timestamp
	)`
	_, err := s.Exec(query)
	return err
}
