package envVar

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

type Vars struct {
	Port         string `env:"PORT"`
	WsPort       string `env:"WS_PORT"`
	DatabaseName string `env:"DATABASE_NAME"`
	DatabaseHost string `env:"DATABASE_HOST"`
	DatabasePort string `env:"DATABASE_PORT"`
	DatabaseUser string `env:"DATABASE_USER"`
	DatabasePass string `env:"DATABASE_PASS"`
	Environment  string `env:"ENVIRONMENT"`
}

var Variables = Vars{}

func Start() {
	options := env.Options{RequiredIfNoDef: true}
	err := env.ParseWithOptions(&Variables, options)

	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}
