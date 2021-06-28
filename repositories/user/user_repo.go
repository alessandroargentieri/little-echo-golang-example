package user

import (
	errmng "templ/errors"
	"templ/models"
	conn "templ/repositories"

	"github.com/labstack/gommon/log"
)

const (
	SAVE_USER_QUERY   = `INSERT INTO users (id, name, surname, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, surname, created_at, updated_at;`
	GET_USER_QUERY    = `SELECT * FROM users where id = $1;`
	UPDATE_USER_QUERY = `UPDATE users SET name = $1, surname = $2, updated_at = $3 WHERE id = $4 RETURNING id, name, surname, created_at, updated_at;`
	DELETE_USER_QUERY = `DELETE FROM users WHERE id = $1;`
)

var Save = func(user models.User) (*models.User, error) {
	var inserted models.User
	err := conn.Db.QueryRow(SAVE_USER_QUERY, user.ID, user.Name, user.Surname, user.CreatedAt, user.UpdatedAt).
		Scan(&inserted.ID, &inserted.Name, &inserted.Surname, &inserted.CreatedAt, &inserted.UpdatedAt)

	if err != nil {
		log.Error(err)
	}
	return &inserted, err
}

var Get = func(id string) (*models.User, error) {
	var user models.User
	err := conn.Db.QueryRow(GET_USER_QUERY, id).
		Scan(&user.ID, &user.Name, &user.Surname, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Error(err)
	}
	return &user, err
}

var Update = func(id string, user models.User) (*models.User, error) {
	if _, err := Get(id); err != nil {
		return nil, errmng.NotFoundError(id)
	}
	var updated models.User
	err := conn.Db.QueryRow(SAVE_USER_QUERY, user.ID, user.Name, user.Surname, user.CreatedAt, user.UpdatedAt).
		Scan(&updated.ID, &updated.Name, &updated.Surname, &updated.CreatedAt, &updated.UpdatedAt)

	if err != nil {
		log.Error(err)
	}
	return &updated, err
}

var Delete = func(id string) error {
	if _, err := conn.Db.Exec(DELETE_USER_QUERY, id); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
