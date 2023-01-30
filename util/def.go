package util

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config *viper.Viper
	Logger *zap.Logger
)

type News struct {
	Zt int        `json:"zt"`
	Tp string     `json:"tp"`
	Lx string     `json:"lx"`
	Lj string     `json:"lj"`
	Wb [][]string `json:"wb"`
}
