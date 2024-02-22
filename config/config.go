package config

import (
	"flag"
	"github.com/iooikaak/frame/cache/redis/v8"
	"github.com/iooikaak/frame/elasticsearch"
	"github.com/iooikaak/frame/gorm"
	"github.com/iooikaak/frame/json"
	rocketmqConfig "github.com/iooikaak/frame/mq/rocketmq/config"
	"github.com/iooikaak/microService2/model/enum"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"io"
	"os"
	"path/filepath"

	"github.com/iooikaak/frame/config"
	"github.com/iooikaak/frame/config/paladin"
	"github.com/iooikaak/frame/config/paladin/apollo"
)

var (
	Conf     = &Configuration{}
	confPath string
	M        paladin.YAML
)

type Configuration struct {
	BaseCfg  *config.BaseCfg   `json:"baseConfig" yaml:"baseConfig"`
	TimeOut  int32             `json:"timeOut" yaml:"timeOut"` //单次处理请求超时时间
	HorseLog string            `json:"horseLog" yaml:"horseLog"`
	IMApi    map[string]string `json:"imApi" yaml:"imApi"`
	//RedisServers     map[string]*RedisServer      `json:"redis" yaml:"redis"`
	MysqlServers     map[string]*MysqlServer        `json:"mysql" yaml:"mysql"`
	TiDBServers      map[string]*TiDBServer         `json:"tiDB" yaml:"tiDB"`
	LiveCfg          *LiveCfg                       `json:"liveCfg" yaml:"liveCfg"`
	EsServers        map[string]*ElasticServer      `json:"esServers" yaml:"esServers"`
	NsqWriters       []string                       `json:"nsqWriters" yaml:"nsqWriters"`
	NsqTopic         string                         `json:"nsqTopic" yaml:"nsqTopic"`
	NsqLookUps       []string                       `json:"nsqLookups" yaml:"nsqLookups"`
	NsqNewCfg        *config.NsqCfg                 `json:"nsqCfg" yaml:"nsqCfg"`
	Elastic          *elasticsearch.ElasticConfig   `json:"elastic" yaml:"elastic"`
	Redis            *redis.Config                  `json:"redis" yaml:"redis"`
	DB               *gorm.Config                   `json:"db" yaml:"db"`
	RocketMq         *rocketmqConfig.RocketmqConfig `json:"rocketMq" yaml:"rocketMq" `
	Hosts            map[string]string              `json:"hosts" yaml:"hosts"`
	Address          string                         `json:"address" yaml:"address"`
	Env              *config.Environment            `json:"env" yaml:"env"`
	WebCfg           *WebCfg                        `json:"webCfg" yaml:"webCfg"`
	RoomSyncInterval int                            `json:"room_sync_interval" yaml:"room_sync_interval"`
	OSSCfg           AliOssCfg                      `json:"ali_oss_cfg" yaml:"ali_oss_cfg"`
	DeadUserDuration int64                          `json:"dead_user_duration" yaml:"dead_user_duration"`
	DefaultOnlineCnt int64                          `json:"default_online_cnt" yaml:"default_online_cnt"`
	Shence           *ShenceCfg                     `json:"shenceCfg" yaml:"shenceCfg"`
	DefaultStickers  []int64                        `json:"default_stickers" yaml:"default_stickers"`
	QiniuSdk         *SdkCfg                        `json:"qiniuSdk" yaml:"qiniuSdk"`
	SnowflakeMachine int64                          `json:"snowflakeMachine" yaml:"snowflakeMachine"`
	Jaeger           string                         `json:"jaeger" yaml:"jaeger"`
}

type RedisServer struct {
	Host             string `json:"host"`
	Password         string `json:"password"`
	RedisIdle        int    `json:"redis_idle"`
	RedisIdleTimeout string `json:"redis_idle_timeout"`
}

type MysqlServer struct {
	Host                 string `json:"host"`
	Database             string `json:"database"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	TimeZone             string `json:"time_zone"`
	MysqlMaxConnections  int    `json:"mysql_max_connections"`
	MysqlIdle            int    `json:"mysql_idle"`
	MysqlConnMaxLifetime int    `json:"mysql_conn_max_lifetime"`
	Charset              string `json:"charset"`
}

type TiDBServer struct {
	Host                 string `json:"host"`
	Database             string `json:"database"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	TimeZone             string `json:"time_zone"`
	MysqlMaxConnections  int    `json:"mysql_max_connections"`
	MysqlIdle            int    `json:"mysql_idle"`
	MysqlConnMaxLifetime int    `json:"mysql_conn_max_lifetime"`
	Charset              string `json:"charset"`
}

type LiveCfg struct {
	AppId               string `json:"appid"`
	Certificate         string `json:"certificate"`
	CustomerId          string `json:"customer_id"` //踢人要用的参数
	CustomerCertificate string `json:"customer_certificate"`
	AgoraAPIUrl         string `json:"agora_api_url"`
	TokenTTL            int64  `json:"token_ttl"`        // 声网token有效时长，秒
	CompareInterval     int    `json:"compare_interval"` // 比较麦位和channel用户差异间隔秒数
	AgoraCdnApiUrl      string `json:"agora_cdn_api_url"`
}

type ElasticServer struct {
	Hosts string `json:"hosts"`
}

type WebCfg struct {
	RedPacketUrl    string `json:"red_packet_url"`
	FollowStatusUrl string `json:"follow_status_url"`
	FollowUrl       string `json:"follow_url"`
	PeiWanUrl       string `json:"pei_wan_url"`
	PeiWanSUrl      string `json:"pei_wans_url"`
	XAppId          string `json:"x_app_id"`
	XAppToken       string `json:"x_app_token"`
	PHPAPIUrl       string `json:"php_api_url"`
}

type AliOssCfg struct {
	Endpoint        string `json:"end_point"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

// 神策埋点统计
type ShenceCfg struct {
	Switch  bool   `json:"switch"`
	Host    string `json:"host"`
	Timeout int    `json:"timeout"`
	Project string `json:"project"`
}

type SdkCfg struct {
	PublishDomain      string `json:"PublishDomain"`
	PlayDomain         string `json:"PlayDomain"`
	HlsDomain          string `json:"HlsDomain"`
	SnapshotDomain     string `json:"SnapshotDomain"`
	Hubname            string `json:"HubName"`
	Accesskey          string `json:"AccessKey"`
	Securitykey        string `json:"SecretKey"`
	ExpireAfterSeconds int64  `json:"ExpireAfterSeconds"`
	Signkey            string `json:"Signkey"`
	SignAvialeTime     int64  `json:"SignAvialeTime"`
}

// apollo配置初始化，依赖本地环境变量
func ApolloInit() error {
	_ = os.Setenv(enum.ApolloNamespaces.String(), enum.ApolloAppName.String())
	_ = os.Setenv(enum.ApolloAppID.String(), enum.ServiceNameLowCase.String())
	_ = os.Setenv(enum.ApolloCluster.String(), enum.ApolloClusterValue.String())
	_ = os.Setenv(enum.ApolloMetaAddr.String(), os.Getenv(enum.ApolloMetaAddrValue.String()))
	_ = os.Setenv(enum.ApolloCacheDir.String(), enum.ApolloCacheDirValue.String())
	if err := paladin.Init(apollo.PaladinDriverApollo); err != nil {
		panic(err)
	}
	if err := paladin.Get(enum.ApolloAppName.String()).UnmarshalJSON(&Conf); err != nil {
		panic(err)
	}
	if err := paladin.Watch(enum.ApolloAppName.String(), &M); err != nil {
		panic(err)
	}
	return nil
}

func CreateTracer(servieName string) (opentracing.Tracer, io.Closer, error) {
	var cfg = jaegercfg.Configuration{
		ServiceName: servieName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: enum.JaegerSampleConfigParam.Float64(),
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			// 按实际情况替换你的 ip
			CollectorEndpoint: Conf.Jaeger,
		},
	}
	jLogger := jaegerlog.StdLogger
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
	)
	return tracer, closer, err
}

func init() {
	flag.StringVar(&confPath, "Conf", "", "conf values")
	//Init()
}

func Init() (err error) {
	var (
		jsonFIle string
	)
	if confPath != "" {
		jsonFIle, err = filepath.Abs(confPath)
	} else {
		jsonFIle, err = filepath.Abs(enum.MicroService2Json.String())
	}
	if err != nil {
		return
	}
	jsonRead, err := os.ReadFile(jsonFIle)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonRead, Conf)
	if err != nil {
		return
	}
	return
}
