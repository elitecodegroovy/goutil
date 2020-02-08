package pkg1

import (
	"github.com/elitecodegroovy/goutil"
)

func init() {
	log.Info("b.go >>>" + goutil.GetCurrentTimeISOStrTime())
}
