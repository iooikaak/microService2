package service

import (
	"github.com/iooikaak/microService2/config"
	show2 "github.com/iooikaak/microService2/dao/es/show"
	"github.com/iooikaak/microService2/dao/redis/show"
	"github.com/iooikaak/microService2/dao/rocketmq/test"
	"sync"

	"github.com/iooikaak/frame/xlog"
)

var (
	srv  *BaseService
	once sync.Once
)

type BaseService struct {
	mq    *test.Dao
	redis *show.Dao
	es    *show2.Dao
}

func New(cfg *config.Configuration) *BaseService {
	once.Do(func() {
		srv = &BaseService{
			mq:    test.New(cfg),
			redis: show.New(cfg),
			es:    show2.New(cfg),
		}
	})
	return srv
}

func (s *BaseService) Stop() {
	err := s.mq.Producer.Stop()
	if err != nil {
		xlog.Errorf("Stop mq failed err: %v", err)
	}
}
