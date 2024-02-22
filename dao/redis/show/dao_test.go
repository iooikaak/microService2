package show

import (
	"context"
	"os"
	"testing"

	conf "github.com/iooikaak/microService2/config"
)

var (
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../build/config.yaml")
	err := conf.ApolloInit()
	if err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}
