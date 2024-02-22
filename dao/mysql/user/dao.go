package user

import (
	conf "github.com/iooikaak/microService2/config"

	"github.com/iooikaak/frame/gorm"
)

const MaxPrice = 999999

type Dao struct {
	Db *gorm.Engine
}

func New(cfg *conf.Configuration) (d *Dao) {
	d = &Dao{}
	d.Db = gorm.New(cfg.DB)
	return
}

func (d *Dao) Stop() error {
	return d.Db.Close()
}
