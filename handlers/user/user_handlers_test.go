package user

import (
	"templ/models"
	service "templ/services/user"
	"templ/utils"
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	_, _ = service.Save(models.User{
		ID:        utils.PointerOfString("ID"),
		Name:      utils.PointerOfString("Name"),
		Surname:   utils.PointerOfString("Surname"),
		CreatedAt: utils.PointerOfTime(time.Now()),
		UpdatedAt: utils.PointerOfTime(time.Now()),
	})
	/*if err != nil {
		t.Errorf(fmt.Sprintf("Test failed! %s", err.Error()))
	}*/
}
