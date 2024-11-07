package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Row struct{}

type RowData struct {
	Id_row  int
	Number  int
	Id_hall int
}

func (r *Row) SelectQuery(query string) []RowData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []RowData
	for rows.Next() {
		var id_row int
		var number int
		var id_hall int

		rows.Scan(&id_row, &number, &id_hall)
		row_table := RowData{
			Id_row:  id_row,
			Number:  number,
			Id_hall: id_hall,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (r *Row) Select() []RowData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Row")
	if err != nil {
		panic(err)
	}
	var array_rows []RowData
	for rows.Next() {
		var id_row int
		var number int
		var id_hall int

		rows.Scan(&id_row, &number, &id_hall)
		row_table := RowData{
			Id_row:  id_row,
			Number:  number,
			Id_hall: id_hall,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (r *Row) Insert(number int, id_hall int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Row(number, id_hall) VALUES(%d, %d)`, number, id_hall)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (r *Row) Update(id_row int, number int, id_hall int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Row SET number = %d, id_hall = %d WHERE id_row = %d`,
		number, id_hall, id_row)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
