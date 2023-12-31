package controller

import "main.go/storage/postgres"

type Controller struct {
	Store postgres.Storage
}
func New(store postgres.Storage) Controller {
	return Controller{
		Store: store,
	}
}