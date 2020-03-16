package api

import (
	"github.com/edwintcloud/print-solution/server/services/print"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// API represents the api
type API struct {
	PrintService *print.Service
	E            *echo.Echo
}

// RegisterHandlers registers api handlers
func (api *API) RegisterHandlers() {
	routes := api.E.Group("/api")
	{
		routes.GET("/clients", api.GetClients)
		routes.POST("/clients", api.RegisterClient)
		routes.POST("/jobs", api.CreatePrintJob)
	}
}

// RegisterMiddlewares registers api middlewares
func (api *API) RegisterMiddlewares() {

	// register request logging middleware
	api.E.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${status} ${method} [${uri}] ${latency_human}\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
	}))

}
