package repository

import (
	"github.com/jmoiron/sqlx"
	restful "github.com/proggcreator/wb-Restful"
)

type EmplWork interface {
	CreateEmpl(employee restful.Employee) (int, error)
	GetAllEmpl() ([]restful.Employee, error)
	GetByIdEmpl(userId int) (restful.Employee, error)
	DeleteEmpl(userId int) error
	UpdateEmpl(userId int, newemployee restful.Employee) error
}
type Repository struct {
	EmplWork
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		EmplWork: NewEmplWorkPostgres(db),
	}
}
