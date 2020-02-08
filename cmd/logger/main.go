package main

import (
	"github.com/elitecodegroovy/goutil"
	_ "github.com/elitecodegroovy/goutil/cmd/logger/pkg"
	_ "github.com/elitecodegroovy/goutil/cmd/logger/pkg1"
	l "github.com/elitecodegroovy/goutil/logger"
	"go.uber.org/zap"
)

var (
	log = l.GetLogger()
)

func init() {
	log.Info("main >>>init ...")
}

func main() {
	log.Info("", zap.String("time", goutil.GetCurrentTimeNumberISOStrTime()))
}
