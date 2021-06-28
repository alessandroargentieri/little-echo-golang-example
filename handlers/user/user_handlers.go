package user

import (
	"net/http"
	errmng "templ/errors"
	"templ/models"
	service "templ/services/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SaveHandler(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, errmng.InvalidInputError())
	}
	saved, err := service.Save(user)
	if err != nil {
		log.Error(err)
		return c.JSON(errmng.ErrResp(err))
	}
	return c.JSON(http.StatusOK, saved)
}

func GetHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error("Invalid id provided in the request")
		return c.JSON(http.StatusBadRequest, errmng.InvalidInputError())
	}
	user, err := service.Get(id)
	if err != nil {
		log.Error(err)
		return c.JSON(errmng.ErrResp(err))
	}
	return c.JSON(http.StatusOK, *user)
}

func UpdateHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error("Invalid id provided in the request")
		return c.JSON(http.StatusBadRequest, errmng.InvalidInputError())
	}
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, errmng.InvalidInputError())
	}
	updated, err := service.Update(id, user)
	if err != nil {
		log.Error(err)
		return c.JSON(errmng.ErrResp(err))
	}
	return c.JSON(http.StatusOK, *updated)
}

func DeleteHandler(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		log.Error("Invalid id provided in the request")
		return c.JSON(http.StatusBadRequest, errmng.InvalidInputError())
	}
	err := service.Delete(id)
	if err != nil {
		log.Error(err)
		return c.JSON(errmng.ErrResp(err))
	}
	return c.JSON(http.StatusOK, "")
}
