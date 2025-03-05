package config

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
)

type Cfg struct {
	Ip string `default:"localhost"`
	Port string `default:"80"`

	TimeAdditionMs int `default:"1"`
	TimeSubtractionMs int `default:"1"`
	TimeMultiplicationsMs int `default:"1"`
	TimeDivisionsMs int `default:"1"`

	ComputingPower int `default:"1"`
}

var Config Cfg

func init() {
	loader := aconfig.LoaderFor(&Config, aconfig.Config{
		Files: []string{"/config.toml", "config.toml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})
	if err := loader.Load(); err != nil {
		panic(err)
	}
	fmt.Println(Config.ComputingPower)
}
