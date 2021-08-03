package app

import (
	"fmt"
	"log"

	"github.com/owen-a183/learngolang.git/config"
	"github.com/owen-a183/learngolang.git/internal/bootstrap"
)

func Run() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("error on load config: %v ", err))
	}
	// fmt.Println(conf)

	bootstrap.InitEcho(conf)
}
