package bootstrap

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/owen-a183/learngolang.git/config"
	echo_adapter "github.com/owen-a183/learngolang.git/internal/adapter/rest/echo"
)

func InitEcho(conf *config.Config) {
	e := echo.New()

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	echoRouter := echo_adapter.EchoRouter{
		Conf:   conf,
		Router: e,
	}

	if err := echoRouter.Start(); err != nil {
		log.Fatal(fmt.Sprintf("failed to start echo: %s", err))
	}
}
