package server

import (
	"github.com/pignuante/go-todoList/api"
	"github.com/pignuante/go-todoList/db"
	"net/http"
)

// Config server config
type Config struct {
	Address     string
	DatabaseURL string
}

// ListenAndServe starts up the server
func ListenAndServe(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Address,
		loggingMiddleware(api.TodoListAPI()))
}
