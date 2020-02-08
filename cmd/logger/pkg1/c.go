package pkg1

import (
	"github.com/elitecodegroovy/goutil"
)

func init() {
	log.Info("c.go >>>" + goutil.GetCurrentTimeISOStrTime())
}
