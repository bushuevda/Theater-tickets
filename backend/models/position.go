package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Position struct{}

type PositionData struct {
	Id_position int
	Full_name   string
	Short_name  string
}

func (p *Position) SelectQuery(query string) []PositionData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []PositionData
	for rows.Next() {
		var id_position int
		var full_name string
		var short_name string

		rows.Scan(&id_position, &full_name, &short_name)
		row_table := PositionData{
			Id_position: id_position,
			Full_name:   full_name,
			Short_name:  short_name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Position) Select() []PositionData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Position")
	if err != nil {
		panic(err)
	}
	var array_rows []PositionData
	for rows.Next() {
		var id_position int
		var full_name string
		var short_name string

		rows.Scan(&id_position, &full_name, &short_name)
		row_table := PositionData{
			Id_position: id_position,
			Full_name:   full_name,
			Short_name:  short_name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Position) Insert(full_name string, short_name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Position(full_name, short_name) VALUES('%s', '%s')`, full_name, short_name)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (p *Position) Update(id_position int, full_name string, short_name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Position SET full_name = '%s', short_name = '%s' WHERE id_position = %d`,
		full_name, short_name, id_position)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
