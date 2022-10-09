package blogs

import (
	"database/sql"
	"errors"
	"time"

	"github.com/LachlanStephan/ls_server/internal/models"
)

type Blog struct {
	Blog_id    int
	User_id    int
	Title      string
	Content    string
	Created_at time.Time
	Updated_at time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(user_id int, title string, content string) (int, error) {
	stmt := `INSERT INTO Blogs (User_id, Title, Content, Created_at, Updated_at) VALUES (?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, user_id, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *BlogModel) Get(id int) (*Blog, error) {
	b := &Blog{}
	stmt := `SELECT Blog_id, User_id, Title, Content, Created_at, Updated_at FROM Blogs WHERE user_id = ?`

	err := m.DB.QueryRow(stmt, id).Scan(&b.Blog_id, &b.User_id, &b.Title, &b.Content, &b.Created_at, &b.Updated_at)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return b, nil
}
