package config

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/theplant/appkit/log"
)

var (
	Config = &AppConfig{}
	Logger = log.Default()
)

type AppConfig struct {
	Port int `default:"4000"`
}

func init() {
	var err error

	if err = configor.New(&configor.Config{ENVPrefix: "ADAPTIVE_LEARNING"}).Load(Config); err != nil {
		Logger.Crit().Log(
			"during", "configor.Load",
			"err", err,
			"msg", fmt.Sprintf("error loading config: %v", err),
		)
		panic(err)
	}

	return
}
