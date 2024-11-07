package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type Ticket struct{}

type TicketData struct {
	Id_ticket              int
	Price                  float64
	Date_performance       string
	Time_begin             string
	Date_sale              string
	Id_performance         int
	Id_physical_person     int
	Id_employment_contract int
	Id_place               int
	Id_measure_unit        int
}

func (t *Ticket) SelectQuery(query string) []TicketData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []TicketData
	for rows.Next() {
		var id_ticket int
		var price float64
		var date_performance string
		var time_begin string
		var date_sale string
		var id_performance int
		var id_physical_person int
		var id_employment_contract int
		var id_place int
		var id_measure_unit int

		rows.Scan(&id_ticket, &price, &date_performance, &time_begin, &date_sale, &id_performance,
			&id_physical_person, &id_employment_contract, &id_place, &id_measure_unit)
		row_table := TicketData{
			Id_ticket:              id_ticket,
			Price:                  price,
			Date_performance:       date_performance,
			Time_begin:             time_begin,
			Date_sale:              date_sale,
			Id_performance:         id_performance,
			Id_physical_person:     id_physical_person,
			Id_employment_contract: id_employment_contract,
			Id_place:               id_place,
			Id_measure_unit:        id_measure_unit,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *Ticket) Select() []TicketData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Ticket")
	if err != nil {
		panic(err)
	}
	var array_rows []TicketData
	for rows.Next() {
		var id_ticket int
		var price float64
		var date_performance string
		var time_begin string
		var date_sale string
		var id_performance int
		var id_physical_person int
		var id_employment_contract int
		var id_place int
		var id_measure_unit int

		rows.Scan(&id_ticket, &price, &date_performance, &time_begin, &date_sale, &id_performance,
			&id_physical_person, &id_employment_contract, &id_place, &id_measure_unit)
		row_table := TicketData{
			Id_ticket:              id_ticket,
			Price:                  price,
			Date_performance:       date_performance,
			Time_begin:             time_begin,
			Date_sale:              date_sale,
			Id_performance:         id_performance,
			Id_physical_person:     id_physical_person,
			Id_employment_contract: id_employment_contract,
			Id_place:               id_place,
			Id_measure_unit:        id_measure_unit,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (t *Ticket) Insert(price float64, date_performance string, time_begin string, date_sale string, id_performance int,
	id_physical_person int, id_employment_contract int, id_place int, id_measure_unit int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Ticket(price, date_performance, time_begin, date_sale, id_performance,
		id_physical_person, id_employment_contract, id_place, id_measure_unit) VALUES(%f, '%s', '%s', '%s', %d, %d, %d, %d, %d)`,
		price, date_performance, time_begin, date_sale, id_performance, id_physical_person,
		id_employment_contract, id_place, id_measure_unit)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (t *Ticket) Update(id_ticket int, price float64, date_performance string, time_begin string, date_sale string, id_performance int,
	id_physical_person int, id_employment_contract int, id_place int, id_measure_unit int) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Ticket SET price = %f, date_performance = '%s', time_begin = '%s', date_sale = '%s',
	id_performance = %d, id_physical_person = %d, id_employment_contract = %d, id_place = %d, id_measure_unit = %d
	WHERE id_ticket = %d`, price, date_performance, time_begin, date_sale, id_performance, id_physical_person,
		id_employment_contract, id_place, id_measure_unit, id_ticket)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
