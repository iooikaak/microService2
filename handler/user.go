package handler

import (
	"github.com/iooikaak/microService2/config"
	"github.com/iooikaak/microService2/service/user"
	pbms2 "github.com/iooikaak/pb/microService2/http"

	frame "github.com/iooikaak/frame/core"
	"github.com/iooikaak/frame/json"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	openLog "github.com/opentracing/opentracing-go/log"
)

var (
	userHandler *UserHandler
)

type UserHandler struct {
	*BaseHandler
	*user.UserService
}

func NewUserHandler(b *BaseHandler) *UserHandler {
	userHandler = &UserHandler{
		BaseHandler: b,
		UserService: user.New(config.Conf),
	}
	return userHandler
}

func (h *BaseHandler) GetUserInfo(c frame.Context) error {
	spanCtx, _ := h.Tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(c.Request().Header))
	span := h.Tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	var req pbms2.GetUserInfoReq
	if err := c.Bind(&req); err != nil {
		return err
	}
	userInfo, err := userHandler.UserService.GetUserInfo(h.Context, &req)
	if err != nil {
		return c.JSON2(1, err.Error(), nil)
	}
	go func() {
		str, _ := json.Marshal(userInfo)
		span.LogFields(
			openLog.String("event", "microService2 GetUserInfo() get data from tidb"),
			openLog.String("value", string(str)),
		)
	}()
	obj := &pbms2.GetUserInfoRep_Data{
		Name: userInfo.Name,
		Age:  userInfo.Age,
		Job:  userInfo.Job,
	}
	if userInfo.Gender == 2 {
		obj.Gender = "女"
	} else {
		//默认男
		obj.Gender = "男"
	}
	go func() {
		b, _ := json.Marshal(obj)
		span.LogFields(
			openLog.String("event", "microService2 GetUserInfo() return data"),
			openLog.String("value", string(b)),
		)
	}()
	return c.JSON2(0, "success", obj)
}
