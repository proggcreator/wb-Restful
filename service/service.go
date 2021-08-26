package service

import (
	"github.com/proggcreator/wb-Restful/repository"
)

type RestfulEmployee interface {
	/*Create(userIdint, employee restful.Employee) (int, error)
	GetAll(userId int) ([]restful.Employee, error)
	GetById(userId int) (restful.Employee, error)
	Delete(userId int) error
	Update(userId int, newemployee restful.Employee) error*/
}
type Service struct {
	RestfulEmployee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
