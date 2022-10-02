package blogs

import (
	"database/sql"
	"time"
)

type Blog struct {
	Blog_id int
	Title string
	Content string
	Created_at time.Time
	Updated_at time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO Blogs (Title, Content, Created_at) VALUES (?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, title, content) 
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}