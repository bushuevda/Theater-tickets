package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Performance struct{}

type PerformanceData struct {
	Id_performance      int
	Name                string
	Duration            float64
	Id_type_performance int
	Id_measure_unit     int
}

func (p *Performance) SelectQuery(query string) []PerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []PerformanceData
	for rows.Next() {
		var id_performance int
		var name string
		var duration float64
		var id_type_performance int
		var id_measure_unit int

		rows.Scan(&id_performance, &name, &duration, &id_type_performance, &id_measure_unit)
		row_table := PerformanceData{
			Id_performance:      id_performance,
			Name:                name,
			Duration:            duration,
			Id_type_performance: id_type_performance,
			Id_measure_unit:     id_measure_unit,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Performance) Select() []PerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Performance")
	if err != nil {
		panic(err)
	}
	var array_rows []PerformanceData
	for rows.Next() {
		var id_performance int
		var name string
		var duration float64
		var id_type_performance int
		var id_measure_unit int

		rows.Scan(&id_performance, &name, &duration, &id_type_performance, &id_measure_unit)
		row_table := PerformanceData{
			Id_performance:      id_performance,
			Name:                name,
			Duration:            duration,
			Id_type_performance: id_type_performance,
			Id_measure_unit:     id_measure_unit,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *Performance) Insert(name string, duration float64, id_type_performance int, id_measure_unit int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Performance(name, duration, id_type_performance, id_measure_unit) VALUES('%s', %f, %d, %d)`,
		name, duration, id_type_performance, id_measure_unit)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (p *Performance) Update(id_performance int, name string, duration float64, id_type_performance int, id_measure_unit int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Performance SET name = '%s', duration = %f, id_type_performance = %d, id_measure_unit = %d
							WHERE id_performance = %d`, name, duration, id_type_performance, id_performance, id_measure_unit)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
