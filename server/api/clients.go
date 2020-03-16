package api

import (
	"net/http"

	"github.com/edwintcloud/print-solution/server/models"
	"github.com/labstack/echo"
)

// GetClients returns a list of clients with printer availability information
func (api *API) GetClients(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}

// RegisterClient registers a new client with the server
func (api *API) RegisterClient(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
