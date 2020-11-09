package logs

import "github.com/astaxie/beego/logs"

var beeLogger *logs.BeeLogger

func SetLogger(logger *logs.BeeLogger) {
	beeLogger = logger
}

func Log() *logs.BeeLogger {
	check()

	return beeLogger
}

func check() {
	if beeLogger == nil {
		panic("beeLog对象未初始化")
	}
}
