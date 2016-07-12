package cron

import (
	//"fmt"
	"log"
	"time"

	"github.com/mesos-utility/elastic-metrics/funcs"
	"github.com/mesos-utility/elastic-metrics/g"
	"github.com/open-falcon/common/model"
)

func Collect() {
	if !g.Config().Transfer.Enable {
		return
	}

	if g.Config().Transfer.Addr == "" {
		return
	}

	go collect(g.Config().Service)
}

func collect(SrvCfg *g.ServiceConfig) {
	// start collect data for elastic cluster.
	for {
		var interval int64 = g.Config().Transfer.Interval
		time.Sleep(time.Duration(interval) * time.Second)

		mvs := []*model.MetricValue{}
		if !SrvCfg.Enable {
			continue
		}

		mvs = funcs.CollectMetrics(SrvCfg)
		if g.Config().Debug {
			log.Printf("%v\n", mvs)
		}

		g.SendToTransfer(mvs)
	}
}
