package log

import (
	"data-monitor/data-common/dispatcher"
	"data-monitor/data-common/models"
	dispatcher2 "data-monitor/data-log/dispatcher"
)

var loggerDispatcher dispatcher.Dispatcher

func AddJob(logDetail *models.LogDetail) {
	loggerDispatcher.AddJob(logDetail)
}

func Start() {
	go func() {
		loggerDispatcher.Start()
	}()
}

func Config(maxWorkerNum, maxJobBufferNum int) {
	if nil != loggerDispatcher {
		loggerDispatcher = dispatcher2.NewLoggerDispatcher(maxWorkerNum, maxJobBufferNum)
	}
}
