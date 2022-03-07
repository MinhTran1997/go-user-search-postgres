package app

import (
	"github.com/core-go/log"
	"github.com/core-go/log/middleware"
	"github.com/core-go/service"
	"github.com/core-go/sql"
)

type Root struct {
	Server     service.ServerConf    `mapstructure:"server"`
	Sql        sql.Config       	 `mapstructure:"sql"`
	Log        log.Config       	 `mapstructure:"log"`
	MiddleWare middleware.LogConfig  `mapstructure:"middleware"`
}