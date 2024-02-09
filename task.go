package main

import "database/sql"

type TaskController interface {
	CreateTask(*Task) error
	DeleteTaskt(int) error
	UpdateTask(*Task) error
	GetTasks() ([]*Task, error)
	GetTaskByID(int) (*Task, error)
}

type PostgresTask struct {
	db *sql.DB
}

func (s *PostgresStore) CreateTaskTable() error {
	query := `create table if not exists task (
        id serial primary key,
        first_name varchar(50),
        last_name varchar(50),
        number serial,
        balance serial,
        created_at timestamp
    )`

	_, err := s.db.Exec(query)
	return err
}
