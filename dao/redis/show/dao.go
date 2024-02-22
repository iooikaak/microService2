package show

import (
	"github.com/iooikaak/microService2/config"

	"github.com/iooikaak/frame/cache/redis/v8"
)

var (
	d *Dao
)

type Dao struct {
	redis *redis.Client
}

//New .
func New(cfg *config.Configuration) (d *Dao) {
	d = &Dao{
		redis: redis.New(cfg.Redis),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() error {
	return d.redis.Close()
}
