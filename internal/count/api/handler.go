package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (srv *Server) GetCount(c echo.Context) error {
	count, err := srv.uc.FetchCount()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]int{"count": count})
}

func (srv *Server) IncreaseCount(c echo.Context) error {
	var input struct {
		Value int `json:"value"`
	}

	if err := c.Bind(&input); err != nil || input.Value == 0 {
		// Если не удалось привязать JSON, пробуем получить из формы
		countStr := c.FormValue("count")
		if countStr == "" {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		value, err := strconv.Atoi(countStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input value")
		}
		input.Value = value
	}

	if err := srv.uc.IncreaseCount(input.Value); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Counter increased by %d", input.Value))
}
