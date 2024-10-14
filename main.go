package main

import (
	"example.com/test/handler"
	"example.com/test/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler := handler.Handler{}
	api := router.API{
		Echo:    e,
		Handler: handler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
