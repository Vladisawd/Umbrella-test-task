package app

import (
	"fmt"
	"log"
	"test1/internal/app/endpoint"
	"test1/internal/app/mw"
	"test1/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e *endpoint.Endpoint

	s *service.Servise

	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
