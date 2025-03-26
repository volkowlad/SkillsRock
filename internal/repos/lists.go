package repos

import (
	sr "SkillsRock"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type ListsDB struct {
	db *pgx.Conn
}

func NewListsDB(db *pgx.Conn) *ListsDB {
	return &ListsDB{db: db}
}

func (l *ListsDB) CreateTask(ctx context.Context, list sr.List) (int, error) {
	tx, err := l.db.Begin(ctx)
	if err != nil {
		slog.Error("ListsDB.CreateTask begin err:", err)
		return -1, err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
INSERT INTO %s (title, description) 
VALUES ($1, $2)
RETURNING id`, listsTable)

	var id int
	err = tx.QueryRow(ctx, query, list.Title, list.Description).Scan(&id)
	if err != nil {
		slog.Error("Error while creating list: ", err.Error())
		return -1, err
	}

	return id, tx.Commit(ctx)
}

func (l *ListsDB) GetList(ctx context.Context) ([]sr.List, error) {
	tx, err := l.db.Begin(ctx)
	if err != nil {
		slog.Error("ListsDB.GetList begin err:", err)
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
SELECT id, title, description, status, created_at, updated_at
FROM %s`, listsTable)

	rows, err := tx.Query(ctx, query)
	if err != nil {
		slog.Error("Error while getting all tasks: ", err.Error())
		return nil, err
	}

	var tasks []sr.List
	for rows.Next() {
		var task sr.List
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			slog.Error("Error while getting one task: ", err.Error())
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Error while getting all tasks: ", err.Error())
		return nil, err
	}

	return tasks, tx.Commit(ctx)
}

func (l *ListsDB) UpdateTask(ctx context.Context, list sr.List, id int) error {
	tx, err := l.db.Begin(ctx)
	if err != nil {
		slog.Error("ListsDB.CreateTask begin err:", err)
		return err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
UPDATE %s
SET description = $1, status = $2, updated_at = now()
WHERE id = $3`, listsTable)

	_, err = tx.Exec(ctx, query, list.Description, list.Status, id)
	if err != nil {
		slog.Error("Error while creating list: ", err.Error())
		return err
	}

	return tx.Commit(ctx)
}

func (l *ListsDB) DeleteTask(ctx context.Context, id int) error {
	tx, err := l.db.Begin(ctx)
	if err != nil {
		slog.Error("ListsDB.CreateTask begin err:", err)
		return err
	}
	defer tx.Rollback(ctx)

	query := fmt.Sprintf(`
DELETE FROM %s
WHERE id = $1`, listsTable)

	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		slog.Error("Error while deleting list: ", err.Error())
		return err
	}

	return tx.Commit(ctx)
}
