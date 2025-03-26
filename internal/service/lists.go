package service

import (
	sr "SkillsRock"
	"SkillsRock/internal/repos"
	"context"
)

type ListsService struct {
	repos *repos.Repos
}

func NewListsService(repos *repos.Repos) *ListsService {
	return &ListsService{
		repos: repos,
	}
}

func (l *ListsService) CreateTask(ctx context.Context, list sr.List) (int, error) {
	return l.repos.CreateTask(ctx, list)
}

func (l *ListsService) GetList(ctx context.Context) ([]sr.List, error) {
	return l.repos.GetList(ctx)
}

func (l *ListsService) UpdateTask(ctx context.Context, list sr.List, id int) error {
	return l.repos.UpdateTask(ctx, list, id)
}

func (l *ListsService) DeleteTask(ctx context.Context, id int) error {
	return l.repos.DeleteTask(ctx, id)
}
