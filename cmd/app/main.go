package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func main() {
	url := "postgres://alexanderp:magic@localhost:5432/alexanderp"
	// this connects & tries a simple 'SELECT 1', panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec("DROP TABLE person;")
	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	err = tx.Commit()

	if err != nil {
		log.Fatalln(err)
	}

	// Selects Mr. Smith from the database
	rows, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "John"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Rows: %v\n", rows)

	var persons []Person
	err = db.Select(&persons, "select * from person")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Persons: %v\n", persons)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))

		if err != nil {
			log.Fatal(err)
		}
	})
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}

}
