package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type MeasureUnit struct{}

type MeasureUnitData struct {
	Id_measure_unit int
	Full_name       string
	Short_name      string
}

func (mu *MeasureUnit) SelectQuery(query string) []MeasureUnitData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []MeasureUnitData
	for rows.Next() {
		var id_measure_unit int
		var full_name string
		var short_name string

		rows.Scan(&id_measure_unit, &full_name, &short_name)
		row_table := MeasureUnitData{
			Id_measure_unit: id_measure_unit,
			Full_name:       full_name,
			Short_name:      short_name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (mu *MeasureUnit) Select() []MeasureUnitData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Measure_unit")
	if err != nil {
		panic(err)
	}
	var array_rows []MeasureUnitData
	for rows.Next() {
		var id_measure_unit int
		var full_name string
		var short_name string

		rows.Scan(&id_measure_unit, &full_name, &short_name)
		row_table := MeasureUnitData{
			Id_measure_unit: id_measure_unit,
			Full_name:       full_name,
			Short_name:      short_name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (mu *MeasureUnit) Insert(full_name string, short_name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf("INSERT INTO Measure_unit(full_name, short_name) VALUES('%s', '%s')", full_name, short_name)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (mu *MeasureUnit) Update(id_measure_unit int, full_name string, short_name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Measure_unit SET full_name = '%s', short_name = '%s' 
							WHERE id_measure_unit = %d`, full_name, short_name, id_measure_unit)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
