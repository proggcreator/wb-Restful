package repository

import "github.com/jmoiron/sqlx"

type RestfulEmployee interface {
	/*Create(userIdint, employee restful.Employee) (int, error)
	GetAll(userId int) ([]restful.Employee, error)
	GetById(userId int) (restful.Employee, error)
	Delete(userId int) error
	Update(userId int, newemployee restful.Employee) error*/
}
type Repository struct {
	RestfulEmployee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
