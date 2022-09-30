package models

import (
	"database/sql"
	"time"
)

type User struct {
	User_id int
	UserName string
	Password string
	IsAdmin bool
	Created_at time.Time
	Updated_at time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(UserName string, Password string, IsAdmin bool) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(User_id int) (*User, error) {
	return nil, nil
}