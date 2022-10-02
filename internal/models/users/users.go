package users

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

// Todo:
// add hashing to password - see if this can be done in the std:: or if we need library (bcrypt) 
func (m *UserModel) Insert(username string, password string, isAdmin bool) (int, error) {
	stmt := `INSERT INTO Users (UserName, Password, IsAdmin, Created_at) VALUES (?, ?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, username, password, isAdmin)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	 
	return int(id), nil
}

func (m *UserModel) Get(User_id int) (*User, error) {
	return nil, nil
}