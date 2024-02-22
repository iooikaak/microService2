package test

import (
	conf "github.com/iooikaak/microService2/config"
	"sync"

	"github.com/iooikaak/frame/mq/rocketmq"
)

var (
	producer *Dao
	once     sync.Once
)

type Dao struct {
	*rocketmq.Producer
}

// New .
func New(cfg *conf.Configuration) (p *Dao) {
	once.Do(func() {
		producer = &Dao{}
		if data, err := rocketmq.NewProducer(cfg.RocketMq); err == nil {
			producer.Producer = data
		} else {
			panic(err)
		}
	})
	return producer
}
