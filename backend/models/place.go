package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Place struct{}

type PlaceData struct {
	Id_place int
	Number   int
	Id_row   int
}

func (p *Place) SelectQuery(query string) []PlaceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []PlaceData
	for rows.Next() {
		var id_place int
		var number int
		var id_row int

		rows.Scan(&id_place, &number, &id_row)
		row_table := PlaceData{
			Id_place: id_place,
			Number:   number,
			Id_row:   id_row,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Place) Select() []PlaceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Place")
	if err != nil {
		panic(err)
	}
	var array_rows []PlaceData
	for rows.Next() {
		var id_place int
		var number int
		var id_row int

		rows.Scan(&id_place, &number, &id_row)
		row_table := PlaceData{
			Id_place: id_place,
			Number:   number,
			Id_row:   id_row,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Place) Insert(number int, id_row int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Place(number, id_row) VALUES(%d, %d)`, number, id_row)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (p *Place) Update(id_place int, number int, id_row int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Place SET number = %d, id_row = %d WHERE id_place = %d`, number, id_row, id_place)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
