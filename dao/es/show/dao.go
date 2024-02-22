package show

import (
	"time"
	//"github.com/iooikaak/frame/xlog"

	v7 "github.com/olivere/elastic/v7"

	"github.com/iooikaak/microService2/config"

	"github.com/iooikaak/frame/elasticsearch"
)

const (
	MaxEsRetries    = 5
	BackOffDelayMin = 100 * time.Millisecond
	BackOffDelayMax = 1 * time.Second
)

type Dao struct {
	es   *v7.Client
	conf *elasticsearch.ElasticConfig
}

func New(cfg *config.Configuration) *Dao {
	instance, err := elasticsearch.NewV7(cfg.Elastic)
	if err != nil {
		panic(err)
	}
	return &Dao{es: instance, conf: cfg.Elastic}
}

type WrapperRetries struct {
	retryNum                         int
	backOffDelayMin, backOffDelayMax time.Duration
}

func (w WrapperRetries) Next(retry int) (time.Duration, bool) {
	if retry > w.retryNum {
		return 0, false
	}
	return backOff(retry, w.backOffDelayMin, w.backOffDelayMax), true
}

func backOff(attempt int, min time.Duration, max time.Duration) time.Duration {
	d := time.Duration(attempt*attempt) * min
	if d > max {
		d = max
	}
	return d
}

func (d *Dao) Stop() {
	d.es.Stop()
}
