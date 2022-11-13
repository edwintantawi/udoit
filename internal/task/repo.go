package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/edwintantawi/udoit/internal/errorx"
	"github.com/edwintantawi/udoit/pkg/idgen"
)

const (
	table    = "tasks"
	errScope = "error task repository"
)

type Repo interface {
	Create(ctx context.Context, v TaskIn) error
	FindAll(ctx context.Context) ([]Task, error)
	FindByID(ctx context.Context, id ID) (Task, error)
	UpdateByID(ctx context.Context, id ID, v TaskIn) error
	DeleteByID(ctx context.Context, id ID) error
}

type repo struct {
	idgen idgen.Generator
	db    *sql.DB
}

func NewRepo(db *sql.DB, idgen idgen.Generator) *repo {
	migration(db)
	return &repo{db: db, idgen: idgen}
}

func migration(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS tasks (
			id			VARCHAR(255)	PRIMARY KEY,
			content		TEXT			NOT NULL,
			description	TEXT			NOT NULL,
			is_done		BOOLEAN 		DEFAULT FALSE,
			created_at	DATETIME		DEFAULT CURRENT_TIMESTAMP,
			updated_at	DATETIME		DEFAULT CURRENT_TIMESTAMP
		)`
	if _, err := db.Exec(q); err != nil {
		panic("tasks table failed to migrate with error: " + err.Error())
	}
}

func (r *repo) Create(ctx context.Context, v TaskIn) error {
	id := ID(fmt.Sprintf("%s-%s", table, r.idgen.NewUUID()))

	q := "INSERT INTO tasks (id, content, description) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, q, id, v.Content, v.Description)
	return err
}

func (r *repo) FindAll(ctx context.Context) ([]Task, error) {
	q := "SELECT id, content, description, is_done, created_at, updated_at FROM tasks"
	row, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	tasks := make([]Task, 0)
	for row.Next() {
		var t Task
		if err := row.Scan(&t.ID, &t.Content, &t.Description, &t.IsDone, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repo) FindByID(ctx context.Context, id ID) (Task, error) {
	q := "SELECT id, content, description, is_done, created_at, updated_at FROM tasks WHERE id = $1"
	var t Task
	row := r.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&t.ID, &t.Content, &t.Description, &t.IsDone, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Task{}, errorx.NewNotFound("Task not found")
		}
		return Task{}, err
	}
	return t, nil
}

func (r *repo) UpdateByID(ctx context.Context, id ID, v TaskIn) error {
	q := "UPDATE tasks SET content = $1, description = $2 WHERE id = $3"
	result, err := r.db.ExecContext(ctx, q, v.Content, v.Description, id)
	if err != nil {
		return err
	}

	c, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if c < 1 {
		return errorx.NewNotFound("Task not found")
	}

	return nil
}

func (r *repo) DeleteByID(ctx context.Context, id ID) error {
	q := "DELETE FROM tasks WHERE id = $1"
	result, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	c, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if c < 1 {
		return errorx.NewNotFound("Task not found")
	}

	return nil
}
