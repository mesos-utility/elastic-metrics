package funcs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mesos-utility/elastic-metrics/g"
	"github.com/open-falcon/common/model"
)

func CollectMetrics(srvCfg *g.ServiceConfig) []*model.MetricValue {
	var interval int64 = g.Config().Transfer.Interval
	hostname, err := g.Hostname()
	if err != nil {
		log.Println("Get hostname failed", err)
	}

	addr := srvCfg.Apiurl
	resp, err := http.Get(addr)

	if err != nil {
		log.Println("get mesos metric data fail", err)
		return nil
	}
	defer resp.Body.Close()

	// read json http response
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read mesos metric fail", err)
		return nil
	}

	var f interface{}
	err = json.Unmarshal(jsonData, &f)
	if err != nil {
		log.Println("Unmarshal metric data fail", err)
		return nil
	}

	now := time.Now().Unix()
	m := f.(map[string]interface{})
	mvs := []*model.MetricValue{}
	for k, v := range m {
		mtype := "GAUGE"
		if t, ok := clusterHealth[k]; ok {
			mtype = t
		}

		key := fmt.Sprintf("elastic.health.%s", k)

		metric := &model.MetricValue{
			Endpoint:  hostname,
			Metric:    key,
			Value:     v,
			Timestamp: now,
			Step:      interval,
			Type:      mtype,
		}

		mvs = append(mvs, metric)
	}

	return mvs
}
