package pkg1

import (
	"github.com/elitecodegroovy/goutil"
	l "github.com/elitecodegroovy/goutil/logger"
)

var (
	log = l.GetLogger()
)

func init() {
	log.Info("a.go >>>" + goutil.GetCurrentTimeISOStrTime())
}
