package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	User_id    int
	FirstName  string
	LastName   string
	Password   []byte
	Email      string
	IsAdmin    bool
	Created_at time.Time
	Updated_at time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(firstName, lastName, email, password string, isAdmin bool) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO Users (FirstName, LastName, Email, Password, IsAdmin, Created_at) VALUES (?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	_, err = m.DB.Exec(stmt, firstName, lastName, email, string(hashedPassword), isAdmin)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "user_email") {
				return ErrDuplicateEmail
			}
			return err
		}
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT User_id, Password FROM Users WHERE Email = ?`

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	exists := false
	stmt := `SELECT EXISTS(SELECT 1 FROM Users WHERE User_id = ?)`
	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

func (m *UserModel) GetAdminUsers() ([]int, error) {
	stmt := `SELECT Users.User_id FROM Users WHERE Users.IsAdmin = 1`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	ids := []int{}
	for rows.Next() {
		id := 0
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
