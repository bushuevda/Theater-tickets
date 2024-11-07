package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type ListPerformance struct{}

type ListPerformanceData struct {
	Id_list_performance int
	Date                string
	Id_theatre          int
	Id_performance      int
}

func (lp *ListPerformance) SelectQuery(query string) []ListPerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []ListPerformanceData
	for rows.Next() {
		var id_list_performance int
		var date string
		var id_theatre int
		var id_performance int

		rows.Scan(&id_list_performance, &date, &id_theatre, &id_performance)
		row_table := ListPerformanceData{
			Id_list_performance: id_list_performance,
			Date:                date,
			Id_theatre:          id_theatre,
			Id_performance:      id_performance,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (lp *ListPerformance) Select() []ListPerformanceData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM List_performance")
	if err != nil {
		panic(err)
	}
	var array_rows []ListPerformanceData
	for rows.Next() {
		var id_list_performance int
		var date string
		var id_theatre int
		var id_performance int

		rows.Scan(&id_list_performance, &date, &id_theatre, &id_performance)
		row_table := ListPerformanceData{
			Id_list_performance: id_list_performance,
			Date:                date,
			Id_theatre:          id_theatre,
			Id_performance:      id_performance,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (lp *ListPerformance) Insert(date string, id_theatre int, id_performance int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf("INSERT INTO List_performance(date, id_theatre, id_performance) VALUES('%s', %d, %d)", date, id_theatre, id_performance)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (lp *ListPerformance) Update(id_list_performace int, date string, id_theatre int, id_performace int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE List_performance SET date = '%s', id_theatre = %d, id_performance = %d
							WHERE id_list_performace = %d`, date, id_theatre, id_performace, id_list_performace)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
