package user

import (
	"flag"
	"os"

	"testing"

	conf "github.com/iooikaak/microService2/config"
)

var srv *UserService

func TestMain(t *testing.M) {
	_ = flag.Set("conf", "../build/config.yaml")
	if err := conf.Init(); err != nil {
		panic(err)
	}
	srv = New(conf.Conf)
	os.Exit(t.Run())
}
