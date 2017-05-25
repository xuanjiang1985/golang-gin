package conf

import (
	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface
var err error

func init() {
	Logger, err = seelog.LoggerFromConfigAsFile("log/seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file ", err)
	}
}
