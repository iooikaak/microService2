package user

import (
	"os"
	"testing"

	conf "github.com/iooikaak/microService2/config"
)

var d *Dao

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../../build/config.yaml")
	if err := conf.ApolloInit(); err != nil {
		panic(err)
	}
	d = New(conf.Conf)
	os.Exit(m.Run())
}
