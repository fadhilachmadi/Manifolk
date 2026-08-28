package repository

import (
	"Manifolk/model"
	"database/sql"
)

type TodoReposPosgre struct {
	db *sql.DB
}

func NewTodoReposPosgre(db *sql.DB) TodoReposPosgre {
	return TodoReposPosgre{
		db: db,
	}
}

func (r *TodoReposPosgre) GetAllData() ([]model.Todo, error) {
	query := `
        SELECT id, name, is_done
        FROM activities
        ORDER BY id
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []model.Todo{}

	for rows.Next() {
		var todo model.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Name,
			&todo.IsDone,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r TodoReposPosgre) CreateData(todo model.Todo) error {
	query := `
        INSERT INTO activities (name, is_done)
        VALUES ($1, $2)
    `

	_, err := r.db.Exec(
		query,
		todo.Name,
		todo.IsDone,
	)

	return err
}

func (r TodoReposPosgre) UpdateDataName(id int, name string) error {
	query := `
        UPDATE activities
        SET name = $1
        WHERE id = $2
    `

	_, err := r.db.Exec(query, name, id)

	return err
}

func (r TodoReposPosgre) UpdateDataStatus(id int, isDone bool) error {
	query := `
        UPDATE activities
        SET is_done = $1
        WHERE id = $2
    `

	_, err := r.db.Exec(query, isDone, id)

	return err
}

func (r TodoReposPosgre) DeletePostgre(id int) error {
	query := `
        DELETE FROM activities
        WHERE id = $1
    `

	_, err := r.db.Exec(query, id)

	return err
}
