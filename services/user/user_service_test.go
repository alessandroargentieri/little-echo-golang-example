package user

import (
	"templ/models"
	repo "templ/repositories/user"
	"templ/utils"
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	repo.Save = func(user models.User) (*models.User, error) {
		return &models.User{}, nil
	}
	_, err := Save(models.User{
		ID:        utils.PointerOfString("ID"),
		Name:      utils.PointerOfString("Name"),
		Surname:   utils.PointerOfString("Surname"),
		CreatedAt: utils.PointerOfTime(time.Now()),
		UpdatedAt: utils.PointerOfTime(time.Now()),
	})
	if err != nil {
		t.Errorf("Test failed!")
	}
}

func TestSaveUser2(t *testing.T) {

	_, err := Save(models.User{
		ID:        utils.PointerOfString("ID"),
		Name:      utils.PointerOfString("Name"),
		Surname:   utils.PointerOfString("Surname"),
		CreatedAt: utils.PointerOfTime(time.Now()),
		UpdatedAt: utils.PointerOfTime(time.Now()),
	})
	if err != nil {
		t.Errorf("Test failed!")
	}
}
