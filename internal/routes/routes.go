package routes

import (
	"duhweb/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", app.ProjectHandler.InitPage)
	r.Get("/about", app.ProjectHandler.AboutPage)
	r.Get("/projects", app.ProjectHandler.ProjectsPage)
	return r
}