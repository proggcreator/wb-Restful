package repository

import (
	"github.com/jmoiron/sqlx"
	restful "github.com/proggcreator/wb-Restful"
)

type EmplWork interface {
	CreateEmpl(empl restful.Employee) (string, error)
	GetAllEmpl() ([]restful.Employee, error)
	GetByIdEmpl(userId string) (restful.Employee, error)
	DeleteEmpl(userId string) error
	UpdateEmpl(newemployee restful.Employee) error
}
type Repository struct {
	EmplWork
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EmplWork: NewEmplWorkPostgres(db),
	}
}
