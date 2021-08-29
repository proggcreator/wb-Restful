package service

import (
	restful "github.com/proggcreator/wb-Restful"
	"github.com/proggcreator/wb-Restful/repository"
)

type EmplWorkService struct {
	repo repository.EmplWork
}

func NewEmplWorkService(repo repository.EmplWork) *EmplWorkService {
	return &EmplWorkService{repo: repo}
}

func (s *EmplWorkService) CreateEmpl(empl restful.Employee) (int, error) {
	return s.repo.CreateEmpl(empl)
}

func (s *EmplWorkService) GetAllEmpl() ([]restful.Employee, error) {
	return s.repo.GetAllEmpl()
}
