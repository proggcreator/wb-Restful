package service

import (
	restful "github.com/proggcreator/wb-Restful"
	"github.com/proggcreator/wb-Restful/repository"
)

type EmplWork struct {
	repo repository.Repository
}

func NewEmplWork(repo repository.Repository) *EmplWork {
	return &EmplWork{repo: repo}
}

func (s *EmplWork) CreateEmpl(empl restful.Employee) (int, error) {
	return s.repo.CreateEmpl(empl)
}

func (s *EmplWork) GetAllEmpl() ([]restful.Employee, error) {
	return s.repo.GetAllEmpl()
}
