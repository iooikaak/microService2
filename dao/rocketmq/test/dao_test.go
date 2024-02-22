package test

import (
	"context"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	conf "github.com/iooikaak/microService2/config"
)

var (
	d        *Dao
	confPath string
	ctx      = context.Background()
)

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../build/config.yaml")
	//err := conf.ApolloInit()
	//if err != nil {
	//	panic(err)
	//}
	Init()
	d = New(conf.Conf)
	os.Exit(m.Run())
}

func Init() (err error) {
	var (
		yamlFile string
	)
	if confPath != "" {
		yamlFile, err = filepath.Abs(confPath)
	} else {
		yamlFile, err = filepath.Abs("../../../config/microService2.yaml")
	}
	if err != nil {
		return
	}
	yamlRead, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlRead, conf.Conf)
	if err != nil {
		return
	}
	return
}
