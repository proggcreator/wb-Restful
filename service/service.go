package service

import (
	restful "github.com/proggcreator/wb-Restful"
	"github.com/proggcreator/wb-Restful/repository"
)

type EmplWork interface {
	CreateEmpl(employee restful.Employee) (int, error)
	GetAllEmpl() ([]restful.Employee, error)
	GetByIdEmpl(userId int) (restful.Employee, error)
	DeleteEmpl(userId int) error
	UpdateEmpl(userId int, newemployee restful.Employee) error
}
type Service struct {
	EmplWork
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		EmplWork: NewEmplWorkService(repos.EmplWork)}
}
