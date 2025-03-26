package service

import (
	sr "SkillsRock"
	"SkillsRock/internal/repos"
	"context"
)

type List interface {
	CreateTask(ctx context.Context, list sr.List) (int, error)
	GetList(ctx context.Context) ([]sr.List, error)
	UpdateTask(ctx context.Context, list sr.List, id int) error
	DeleteTask(ctx context.Context, id int) error
}

type Service struct {
	List
}

func NewService(repos *repos.Repos) *Service {
	return &Service{
		List: NewListsService(repos),
	}
}
