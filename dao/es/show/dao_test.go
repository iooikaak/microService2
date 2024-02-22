package show

import (
	"context"
	"os"
	"testing"

	"github.com/iooikaak/microService2/config"
)

var (
	d   *Dao
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../build/config.yaml")
	err := config.ApolloInit()
	if err != nil {
		panic(err)
	}
	d = New(config.Conf)
	os.Exit(m.Run())
}
