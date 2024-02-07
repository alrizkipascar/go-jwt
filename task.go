package main

type Task interface {
	CreateTask(*Account) error
	DeleteTaskt(int) error
	UpdateTask(*Account) error
	GetTasks() ([]*Account, error)
	GetTaskByID(int) (*Account, error)
}
