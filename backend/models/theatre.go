package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Theatre struct{}

type TheatreData struct {
	Id_theatre  int
	Name        string
	Id_locality int
}

func (t *Theatre) SelectQuery(query string) []TheatreData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []TheatreData
	for rows.Next() {
		var id_theatre int
		var name string
		var id_locality int

		rows.Scan(&id_theatre, &name, &id_locality)
		row_table := TheatreData{
			Id_theatre:  id_theatre,
			Name:        name,
			Id_locality: id_locality,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *Theatre) Select() []TheatreData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Theatre")
	if err != nil {
		panic(err)
	}
	var array_rows []TheatreData
	for rows.Next() {
		var id_theatre int
		var name string
		var id_locality int

		rows.Scan(&id_theatre, &name, &id_locality)
		row_table := TheatreData{
			Id_theatre:  id_theatre,
			Name:        name,
			Id_locality: id_locality,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *Theatre) Insert(name string, id_locality int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Theatre(name, id_locality) VALUES('%s', %d)`, name, id_locality)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (t *Theatre) Update(id_theatre int, name string, id_locality int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Theatre SET name = '%s', id_locality = %d WHERE id_theatre = %d`,
		name, id_locality, id_theatre)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
