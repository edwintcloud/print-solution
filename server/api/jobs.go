package api

import (
	"net/http"

	"github.com/edwintcloud/print-solution/server/models"
	"github.com/labstack/echo"
)

// CreatePrintJob creates a new print job
func (api *API) CreatePrintJob(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Response{
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
