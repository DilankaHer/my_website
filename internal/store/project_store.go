package store

import "database/sql"

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	TechStack   string `json:"tech_stack"`
}

type SQLiteProjectStore struct {
	db *sql.DB
}

func NewSQLiteProjectStore(db *sql.DB) *SQLiteProjectStore {
	return &SQLiteProjectStore{db: db}
}

func (s *SQLiteProjectStore) GetAllProjects() ([]Project, error) {
	rows, err := s.db.Query("SELECT id, title, description, link, tech_stack FROM projects")	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Link, &p.TechStack); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}