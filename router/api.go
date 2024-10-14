package router

import (
	"example.com/test/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo    *echo.Echo
	Handler handler.Handler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/year", api.Handler.HandlerSum)
}
