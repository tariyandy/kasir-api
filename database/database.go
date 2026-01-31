package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25) //artinya mendefine maks open conneksi di database kita.
	db.SetMaxIdleConns(5)  //jika tidak ada transaksi, kita menyediakan 5 slot yang tersedia

	log.Println("Database connected successfully")
	return db, nil
}
