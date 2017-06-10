package conf

import (
	seelog "github.com/cihub/seelog"
	"os"
)

var Logger seelog.LoggerInterface
var err error

var basePath = os.Getenv("GOPATH")

func init() {
	Logger, err = seelog.LoggerFromConfigAsFile(basePath + "/src/golang-gin/log/seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file ", err)
	}
}
