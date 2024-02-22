package enum

type Gateway string

const (
	ServiceName              Gateway = "microService2"
	ServiceNameLowCase       Gateway = "microservice2"
	MicroService2Json        Gateway = "config/microService2.json"
	MicroService2JsonDaoPath Gateway = "../../../../config/microService2.json"
)

func (g Gateway) String() string {
	return string(g)
}
