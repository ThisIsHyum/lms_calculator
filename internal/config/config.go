package config

import (
	"flag"
)

var Ip = flag.String("ip", "", "ip of api")
var Port = flag.String("port", "80", "port of api")

func init() {
	flag.Parse()
}
