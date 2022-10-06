package fixture

import (
	"context"
	"database/sql"
	"io/fs"
	"log"

	"github.com/maragudk/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func Up(migratoinDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	if err := migrate.Up(context.Background(), db, migratoinDir); err != nil {
		panic(err)
	}
	return db
}

func Down(db *sql.DB, migratoinDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migratoinDir); err != nil {
		panic(err)
	}
	db.Close()
}
