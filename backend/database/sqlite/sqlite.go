package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const Name_db string = "theatre.db"

type SQlite struct {
	Name_db string
}

func (s *SQlite) Connect() *sql.DB {
	db, err := sql.Open("sqlite3", s.Name_db)
	if err != nil {
		panic(err)
	}
	return db
}

func (s *SQlite) Close(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	return err
}

func (s *SQlite) Select(query string) (*sql.Rows, error) {
	db := s.Connect()

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	s.Close(db)
	return rows, err
}

func (s *SQlite) Insert(query string) error {
	db := s.Connect()
	_, err := db.Exec(query)
	s.Close(db)
	return err
}

func (s *SQlite) Update(query string) error {
	db := s.Connect()
	_, err := db.Exec(query)
	s.Close(db)
	return err
}
