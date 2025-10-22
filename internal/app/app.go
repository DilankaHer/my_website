package app

import (
	"database/sql"
	"duhweb/internal/api"
	"duhweb/internal/store"
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
	ProjectHandler *api.ProjectHandler
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	sqlDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	ProjectHandler := api.NewProjectHandler(sqlDB)
	app := &Application{
		Logger: logger,
		ProjectHandler: ProjectHandler,
		DB: sqlDB,
	}
	return app, nil
}