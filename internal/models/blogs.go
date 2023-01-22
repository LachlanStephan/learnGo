package models

import (
	"database/sql"
	"errors"
	"html/template"
	"time"
)

type Blog struct {
	Blog_id       int
	User_id       int
	FirstName     string
	LastName      string
	Title         string
	Content       template.HTML
	Created_at    time.Time
	Updated_at    time.Time
	FormattedDate string
}

type BlogLink struct {
	Blog_id   int
	Title     string
	FirstName string
	LastName  string
}

type BlogModel struct {
	DB *sql.DB
}

// for now can hardcode the user id as mine
// do this until there is some auth added
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

func (m *BlogModel) GetAll() ([]*BlogLink, error) {
	stmt := `SELECT blogs.Blog_id, blogs.Title, users.FirstName, users.LastName FROM blogs JOIN Users ON Users.user_id = blogs.user_id ORDER BY blogs.Created_at DESC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := []*BlogLink{}

	for rows.Next() {
		b := &BlogLink{}

		err = rows.Scan(&b.Blog_id, &b.Title, &b.FirstName, &b.LastName)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (m *BlogModel) Recent() ([]*BlogLink, error) {
	stmt := `SELECT blogs.Blog_id, blogs.Title, users.FirstName, users.LastName FROM blogs JOIN Users ON Users.user_id = blogs.user_id ORDER BY blogs.Created_at DESC LIMIT 3`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := []*BlogLink{}

	for rows.Next() {
		b := &BlogLink{}

		err = rows.Scan(&b.Blog_id, &b.Title, &b.FirstName, &b.LastName)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (m *BlogModel) Get(id int) (*Blog, error) {
	b := &Blog{}
	stmt := `SELECT Blogs.Blog_id, Blogs.User_id, Blogs.Title, Blogs.Content, Blogs.Created_at, Blogs.Updated_at, Users.FirstName, Users.LastName FROM blogs JOIN Users ON Blogs.User_id = Users.User_id WHERE Blog_id = ?`

	err := m.DB.QueryRow(stmt, id).Scan(&b.Blog_id, &b.User_id, &b.Title, &b.Content, &b.Created_at, &b.Updated_at, &b.FirstName, &b.LastName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return b, nil
}
