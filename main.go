package main

import (
	"context"
	"flag"
	"net/http"
	"sync"

	"github.com/iooikaak/frame/stat/metric/metrics"
	"github.com/iooikaak/frame/xlog"
	"github.com/iooikaak/microService2/model/enum"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/opentracing/opentracing-go"

	pb "github.com/iooikaak/pb/microService2/http"

	"github.com/iooikaak/microService2/config"
	"github.com/iooikaak/microService2/handler"
)

var (
	cfgFile  = flag.String("config-path", "/config/microService2.yaml", "config file")
	logLevel = flag.String("level", "debug", "log level")
	logPath  = flag.String("log-path", "/log/microService2.log", "log path")
)

func main() {
	//flag.Parse()
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	panic(err)
	//}
	//os.Chdir(dir)
	//log.SetLevel(*logLevel)
	//log.SetLogPaths([]string{dir + *logPath})
	////配置文件反序列化到结构体里
	//var cfg = &config.Configuration{}
	//config.Parseconfig(dir+*cfgFile, cfg)

	//远程读取apollo配置文件
	//初始化apollo
	if err := config.ApolloInit(); err != nil {
		panic(err)
	}
	//开启jaeger追踪
	tracer, closer, err := config.CreateTracer(enum.ServiceName.String())
	defer closer.Close()
	if err != nil {
		panic(err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan(enum.JaegerStartSpan.String())
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	//metrics
	go func() {
		http.Handle("/microService2/metrics", promhttp.Handler())
		var m *metrics.Metrics
		var h http.Handler
		var once = sync.Once{}
		once.Do(
			func() {
				m = metrics.New(enum.ServiceName.String())
			},
		)
		once.Do(
			func() {
				h = promhttp.InstrumentMetricHandler(m.RegInstance(), promhttp.HandlerFor(m.Gather(), promhttp.HandlerOpts{}))
			},
		)
		//h := promhttp.InstrumentMetricHandler(m.RegInstance(), promhttp.HandlerFor(m.Gather(), promhttp.HandlerOpts{}))
		err = http.ListenAndServe(":3203", h)
		if err != nil {
			xlog.Fatal("ListenAndServe: ", err.Error())
			return
		}
	}()

	//初始化handler
	server := handler.NewBaseHandler(ctx, tracer)
	//注册到frame
	pb.RegisterMicroService2Server(server, config.Conf.BaseCfg)
}
