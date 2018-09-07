package main

import (
	"github.com/fancygo/fc_log"
	_ "github.com/fancygo/fc_util"
	"time"
)

func main() {
	log, err := fc_log.NewLogger("test", fc_log.LV_TRACE, fc_log.LOG_OUTPUT_SF)
	if err != nil {
		log.Default("new logger err = %v\n", err)
		return
	}
	start := time.Now()
	for i := 0; i < 10; i++ {
		log.Trace("fancy1")
		log.Info("fancy2")
		log.Debug("fancy3")
		log.Warn("fancy4")
		log.Err("fancy5")
		log.Fatal("fancy6")
	}
	secs := time.Since(start).Seconds()
	log.Trace("%v", secs)
}
