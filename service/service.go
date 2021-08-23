package service

import restful "github.com/proggcreator/wb-Restful"

type Service struct {
	RestfulEmployee
}
type RestfulEmployee interface {
	Create(userIdint, employee restful.Employee) (int, error)
	GetAll(userId int) ([]restful.Employee, error)
	GetById(userId int) (restful.Employee, error)
	Delete(userId int) error
	Update(userId int, newemployee restful.Employee) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		RestfulEmployee: NewRestfulService(repos.Employee, repos.TodoList),
	}
}
