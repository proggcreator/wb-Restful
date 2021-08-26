package repository

import (
	"github.com/jmoiron/sqlx"
	restful "github.com/proggcreator/wb-Restful"
)

type EmplWorkPostgres struct {
	db *sqlx.DB
}

func NewEmplWorkPostgres(db *sqlx.DB) *EmplWorkPostgres {
	return &EmplWorkPostgres{db: db}
}

func (s *EmplWorkPostgres) CreateEmpl(empl restful.Employee) (int, error) {
	return 0, nil
}
