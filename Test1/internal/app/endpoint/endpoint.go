package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()

	message := fmt.Sprintf("Days Left: %d", d)

	err := ctx.JSON(http.StatusOK, message)

	if err != nil {
		return err
	}

	return err
}
