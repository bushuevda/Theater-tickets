package models

import (
	"fmt"
	"theatre/database/sqlite"
)

type PhysicalPerson struct{}

type PhysicalPersonData struct {
	Id_physical_person int
	Surname            string
	Name               string
	Patronymic         string
	Date_birthday      string
}

func (p *PhysicalPerson) SelectQuery(query string) []PhysicalPersonData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select(query)
	if err != nil {
		panic(err)
	}
	var array_rows []PhysicalPersonData
	for rows.Next() {
		var id_physical_person int
		var surname string
		var name string
		var patronymic string
		var date_birthday string

		rows.Scan(&id_physical_person, &surname, &name, &patronymic, &date_birthday)
		row_table := PhysicalPersonData{
			Id_physical_person: id_physical_person,
			Surname:            surname,
			Name:               name,
			Patronymic:         patronymic,
			Date_birthday:      date_birthday,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *PhysicalPerson) Select() []PhysicalPersonData {

	db := sqlite.SQlite{Name_db: sqlite.Name_db}
	rows, err := db.Select("SELECT * FROM Physical_person")
	if err != nil {
		panic(err)
	}
	var array_rows []PhysicalPersonData
	for rows.Next() {
		var id_physical_person int
		var surname string
		var name string
		var patronymic string
		var date_birthday string

		rows.Scan(&id_physical_person, &surname, &name, &patronymic, &date_birthday)
		row_table := PhysicalPersonData{
			Id_physical_person: id_physical_person,
			Surname:            surname,
			Name:               name,
			Patronymic:         patronymic,
			Date_birthday:      date_birthday,
		}
		array_rows = append(array_rows, row_table)
	}
	return array_rows
}

func (p *PhysicalPerson) Insert(surname string, name string, patronymic string, date_birthday string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`INSERT INTO Physical_person(surname, name, patronymic, date_birthday) VALUES('%s', '%s', '%s', '%s')`,
		surname, name, patronymic, date_birthday)
	err := db.Insert(query)
	if err != nil {
		panic(err)
	}
}

func (p *PhysicalPerson) Update(id_physical_person int, surname string, name string, patronymic string, date_birthday string) {
	db := sqlite.SQlite{Name_db: sqlite.Name_db}

	query := fmt.Sprintf(`UPDATE Physical_person SET surname = '%s', name = '%s', patronymic = '%s',  date_birthday = '%s'
							WHERE id_physical_person = %d`, surname, name, patronymic, date_birthday, id_physical_person)
	err := db.Update(query)
	if err != nil {
		panic(err)
	}
}
