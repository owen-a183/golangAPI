package echo

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/owen-a183/learngolang.git/config"
)

type EchoRouter struct {
	Conf   *config.Config
	Router *echo.Echo
}

func (e *EchoRouter) Start() error {
	return e.Router.Start(fmt.Sprintf(": %d", e.Conf.Port))
}
