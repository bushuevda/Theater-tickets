package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type EmploymentContract struct{}

type EmploymentContractData struct {
	Id_employment_contract int
	Date_begin             string
	Date_end               string
	Salary                 float64
	Id_physical_person     int
	Id_position            int
	Id_theatre             int
}

func (ec *EmploymentContract) SelectQuery(query string) []EmploymentContractData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []EmploymentContractData
	for rows.Next() {
		var id_employment_contract int
		var date_begin string
		var date_end string
		var salary float64
		var id_physical_person int
		var id_position int
		var id_theatre int

		rows.Scan(&id_employment_contract, &date_begin, &date_end, &salary, &id_physical_person, &id_position, &id_theatre)
		row_table := EmploymentContractData{
			Id_employment_contract: id_employment_contract,
			Date_begin:             date_begin,
			Date_end:               date_end,
			Salary:                 salary,
			Id_physical_person:     id_physical_person,
			Id_position:            id_position,
			Id_theatre:             id_theatre,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (ec *EmploymentContract) Select() []EmploymentContractData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Employment_contract")
	if err != nil {
		panic(err)
	}
	var array_rows []EmploymentContractData
	for rows.Next() {
		var id_employment_contract int
		var date_begin string
		var date_end string
		var salary float64
		var id_physical_person int
		var id_position int
		var id_theatre int

		rows.Scan(&id_employment_contract, &date_begin, &date_end, &salary, &id_physical_person, &id_position, &id_theatre)
		row_table := EmploymentContractData{
			Id_employment_contract: id_employment_contract,
			Date_begin:             date_begin,
			Date_end:               date_end,
			Salary:                 salary,
			Id_physical_person:     id_physical_person,
			Id_position:            id_position,
			Id_theatre:             id_theatre,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (ec *EmploymentContract) Insert(date_begin string, date_end string, salary float64, id_physical_person int, id_position int, id_theatre int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Employment_contract(date_begin, date_end, salary, id_physical_person, id_position, id_theatre)
							 VALUES('%s', '%s', %f, %d, %d, %d)`, date_begin, date_end, salary, id_physical_person, id_position, id_theatre)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (ec *EmploymentContract) Update(id_employment_contract int, date_begin string, date_end string,
	salary float64, id_physical_person int, id_position int, id_theatre int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Employment_contract SET date_begin = '%s', date_end = '%s', salary =  %f, 
		id_physical_person = %d, id_position = %d, id_theatre = %d WHERE id_employment_contract = %d`,
		date_begin, date_end, salary, id_physical_person, id_position, id_theatre, id_employment_contract)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
