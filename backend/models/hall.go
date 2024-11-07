package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Hall struct{}

type HallData struct {
	Id_hall    int
	Name       string
	Id_theatre int
}

func (h *Hall) SelectQuery(query string) []HallData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []HallData
	for rows.Next() {
		var id_hall int
		var name string
		var id_theatre int

		rows.Scan(&id_hall, &name, &id_theatre)
		row_table := HallData{
			Id_hall:    id_hall,
			Name:       name,
			Id_theatre: id_theatre,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (h *Hall) Select() []HallData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Hall")
	if err != nil {
		panic(err)
	}
	var array_rows []HallData
	for rows.Next() {
		var id_hall int
		var name string
		var id_theatre int

		rows.Scan(&id_hall, &name, &id_theatre)
		row_table := HallData{
			Id_hall:    id_hall,
			Name:       name,
			Id_theatre: id_theatre,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (h *Hall) Insert(name string, id_theatre int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf("INSERT INTO Hall(name, id_theatre) VALUES('%s', %d)", name, id_theatre)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (h *Hall) Update(id_hall int, name string, id_theatre int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Hall SET name = '%s', id_theatre = %d 
							WHERE id_hall = %d`, name, id_theatre, id_hall)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
