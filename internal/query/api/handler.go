package api

import (
	"net/http"

	"web-10/internal/query/model"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Parameter 'name' is required")
	}

	user, err := srv.uc.GetUser(name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (srv *Server) AddUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON format")
	}

	if err := srv.uc.AddUser(user.Name); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, "User added successfully")
}
