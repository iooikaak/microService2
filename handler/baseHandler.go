package handler

import (
	"context"
	"sync"

	"github.com/iooikaak/microService2/config"
	"github.com/iooikaak/microService2/service"
	"github.com/opentracing/opentracing-go"
)

var (
	instance *BaseHandler
	once     sync.Once
)

type BaseHandler struct {
	*service.BaseService
	context.Context
	opentracing.Tracer
	*UserHandler
}

func NewBaseHandler(ctx context.Context, tracer opentracing.Tracer) *BaseHandler {
	once.Do(func() {
		instance = &BaseHandler{
			BaseService: service.New(config.Conf),
			Context:     ctx,
			Tracer:      tracer,
		}
		//用户相关的路由方法
		instance.UserHandler = NewUserHandler(instance)
		//xxx相关的路由方法
		//xxx相关的路由方法
	})
	return instance
}
