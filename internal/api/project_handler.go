package api

import (
	"database/sql"
	"duhweb/internal/store"
	"html/template"
	"net/http"
)

type ProjectHandler struct {
	Templates  *template.Template
	ProjectStore *store.SQLiteProjectStore
}

func (h *ProjectHandler) Render(w http.ResponseWriter, tmpl string, data interface{}) {
	if err := h.Templates.ExecuteTemplate(w, tmpl, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewProjectHandler(db *sql.DB) *ProjectHandler {
	return &ProjectHandler{
		Templates:  template.Must(template.ParseGlob("views/*.html")),
		ProjectStore: store.NewSQLiteProjectStore(db),
	}
}

func (h *ProjectHandler) InitPage(w http.ResponseWriter, r *http.Request) {
	h.Render(w, "index", nil)
}

func (h *ProjectHandler) AboutPage(w http.ResponseWriter, r *http.Request) {
	h.Render(w, "about", nil)
}

func (h *ProjectHandler) ProjectsPage(w http.ResponseWriter, r *http.Request) {
	projects, err := h.ProjectStore.GetAllProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Render(w, "projects", projects)
}