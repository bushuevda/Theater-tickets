package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Locality struct{}

type LocalityData struct {
	Id_locality int
	Name        string
}

func (l *Locality) SelectQuery(query string) []LocalityData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []LocalityData
	for rows.Next() {
		var id_locality int
		var name string

		rows.Scan(&id_locality, &name)
		row_table := LocalityData{
			Id_locality: id_locality,
			Name:        name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (l *Locality) Select() []LocalityData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Locality")
	if err != nil {
		panic(err)
	}
	var array_rows []LocalityData
	for rows.Next() {
		var id_locality int
		var name string

		rows.Scan(&id_locality, &name)
		row_table := LocalityData{
			Id_locality: id_locality,
			Name:        name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (l *Locality) Insert(name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf("INSERT INTO Locality(name) VALUES('%s')", name)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (l *Locality) Update(id_locality int, name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Locality SET name = '%s'
							WHERE id_locality = %d`, name, id_locality)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
