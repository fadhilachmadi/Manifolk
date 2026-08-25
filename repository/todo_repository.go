package repository

import (
	"Manifolk/model"
	"database/sql"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return TodoRepository{
		db: db,
	}
}

func (r TodoRepository) GetAll() ([]model.Todo, error) {
	query := `
		SELECT id, name, is_done
		FROM todos
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var todos []model.Todo

	for rows.Next() {
		var todo model.Todo

		if err := rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.IsDone,
		); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r TodoRepository) Create(todo *model.Todo) error {

	query := `
		INSERT INTO todos (name, is_done)
		VALUES (?, ?)
	`

	result, err := r.db.Exec(
		query,
		todo.Name,
		todo.IsDone,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = int(id)

	return nil
}

func (r TodoRepository) Delete(id int) error {
	query := `
		DELETE FROM todos
		WHERE id = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r TodoRepository) UpdateStatus(id int, isDone bool) error {
	query := `
		UPDATE todos
		SET is_done = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, isDone, id)

	if err != nil {
		return err
	}

	return nil
}

func (r TodoRepository) UpdateName(id int, name string) error {
	query := `
		UPDATE todos
		SET name = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, name, id)
	if err != nil {
		return err
	}

	return nil
}
