package repos

import (
	sr "SkillsRock"
	"context"
	"github.com/jackc/pgx/v5"
)

type List interface {
	CreateTask(ctx context.Context, list sr.List) (int, error)
	GetList(ctx context.Context) ([]sr.List, error)
	UpdateTask(ctx context.Context, list sr.List, id int) error
	DeleteTask(ctx context.Context, id int) error
}

type Repos struct {
	List
}

func NewRepos(db *pgx.Conn) *Repos {
	return &Repos{
		List: NewListsDB(db),
	}
}
