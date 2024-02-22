package user

import (
	conf "github.com/iooikaak/microService2/config"
	"github.com/iooikaak/microService2/dao/mysql/user"
	"github.com/iooikaak/microService2/service"
)

type UserService struct {
	*service.BaseService
	db *user.Dao
}

func New(cfg *conf.Configuration) *UserService {
	srv := &UserService{
		BaseService: service.New(cfg),
		db:          user.New(cfg),
	}
	return srv
}
