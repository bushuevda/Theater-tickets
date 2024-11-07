package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type TypePerformance struct{}

type TypePerformanceData struct {
	Id_type_performance int
	Name                string
}

func (t *TypePerformance) SelectQuery(query string) []TypePerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []TypePerformanceData
	for rows.Next() {
		var id_type_performance int
		var name string

		rows.Scan(&id_type_performance, &name)
		row_table := TypePerformanceData{
			Id_type_performance: id_type_performance,
			Name:                name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *TypePerformance) Select() []TypePerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Type_performance")
	if err != nil {
		panic(err)
	}
	var array_rows []TypePerformanceData
	for rows.Next() {
		var id_type_performance int
		var name string

		rows.Scan(&id_type_performance, &name)
		row_table := TypePerformanceData{
			Id_type_performance: id_type_performance,
			Name:                name,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *TypePerformance) Insert(name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Type_performance(name) VALUES('%s')`, name)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (t *TypePerformance) Update(id_type_performance int, name string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Type_performance SET name = '%s'WHERE id_type_performance = %d`,
		name, id_type_performance)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
