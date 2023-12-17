package models

import (
	"github.com/robertjbass/go-rest-api/db"
	"github.com/robertjbass/go-rest-api/utils"
)

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password)
	VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id
	return err
}
