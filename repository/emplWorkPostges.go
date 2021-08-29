package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	restful "github.com/proggcreator/wb-Restful"
)

type EmplWorkPostgres struct {
	db *sqlx.DB
}

func NewEmplWorkPostgres(db *sqlx.DB) *EmplWorkPostgres {
	return &EmplWorkPostgres{db: db}
}

func (s *EmplWorkPostgres) CreateEmpl(empl restful.Employee) error {

	rows, err := s.db.Query("SELECT * FROM employees.employee_add ($1,$2,$3,$4,$5,$6,$7)",
		empl.Id, empl.Name, empl.Last_name, empl.Patronymic,
		empl.Phone, empl.Position, empl.Good_job_count)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (s *EmplWorkPostgres) GetAllEmpl() ([]restful.Employee, error) {
	var lists []restful.Employee
	query := fmt.Sprintf("SELECT * FROM employees. get_all ()")
	err := s.db.Select(&lists, query)
	if err != nil {
		return nil, err
	}

	return lists, nil
}
