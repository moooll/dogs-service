package storage

import (
	"errors"

	"github.com/google/uuid"
	er "github.com/moooll/dogs-service/internal/errors"
	"github.com/moooll/dogs-service/internal/models"
)

// Register registers new user to the storage
func (st *Storage) Register(username, password string) (uuid.UUID, error) {
	id := uuid.New()
	row, err := st.conn.Query(st.ctx, "insert into users(id, username, password) values ($1, $2, $3)",
		id, username, password)
	if err != nil {
		return uuid.Nil, err
	}

	row.Close()

	return id, nil
}

// Check checks is the user is present in the storage (if not, returns er.AuthErrNotRegistered error)
// and if the password provided is the right one (if not, returns er.AuthErrWrongPwd )
func (st *Storage) Check(username, password string) error {
	pw := ""
	err := st.conn.QueryRow(st.ctx, "select password from users where username = $1", username).Scan(&pw)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return errors.New(er.AuthErrNotRegistered)
		}

		return err
	}

	if pw != password {
		return errors.New(er.AuthErrWrongPwd)
	}

	return nil
}

// GetUser retrieves user by id from the database
func (st *Storage) GetUser(id uuid.UUID) (models.User, error) {
	u := models.User{}
	err := st.conn.QueryRow(st.ctx, "select * from users where id = $1", id).Scan(u.ID, u.Username, u.Password)
	if err != nil {
		return models.User{}, nil
	}

	return u, nil
}

// GetUserByName retrieves user by username from the database
func (st *Storage) GetUserByName(name string) (id uuid.UUID, err error) {
	u := models.User{}
	err = st.conn.QueryRow(st.ctx, "select *from users where username = $1", name).Scan(&u.ID, &u.Username, &u.Password)
	if err != nil {
		return uuid.Nil, err
	}

	return u.ID, nil
}