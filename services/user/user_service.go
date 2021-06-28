package user

import (
	"templ/models"
	repo "templ/repositories/user"
)

var Save = func(user models.User) (*models.User, error) {
	return repo.Save(user)
}

var Get = func(id string) (*models.User, error) {
	return repo.Get(id)
}

var Update = func(id string, user models.User) (*models.User, error) {
	return repo.Update(id, user)
}

var Delete = func(id string) error {
	return repo.Delete(id)
}
