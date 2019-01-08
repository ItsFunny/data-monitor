package dispatcher

import (
	"encoding/json"
	"fmt"
	"log"

	"data-monitor/data-common/dispatcher"
	"data-monitor/data-common/models"
	"data-monitor/data-common/worker"
)

// FIXME  这个工程的存在有待商榷

var loggerConsumeFunc = func(value interface{}) error {
	logModel := &models.Logger{}
	if bytes, ok := value.([]byte); ok {
		marshalError := json.Unmarshal(bytes, logModel)
		if nil != marshalError {
			// 2019-01-08 15:04 TODO 持久化到文件中
			log.Println(marshalError)
		} else {
			fmt.Println(logModel)
		}
	}
	return nil
}

type LoggerDispatcher struct {
	dispatcher.BaseDispatcher
}

// 内部类
type loggerWorker struct {
	worker.BaseWoker
}

// 构造返回Dispatcher
func NewLoggerDispatcher(maxWorker, maxBuffer int) dispatcher.Dispatcher {
	dispatcher := &LoggerDispatcher{
		BaseDispatcher: dispatcher.BaseDispatcher{
			MaxWorker: maxWorker,
			WorkQueue: make(chan chan interface{}, maxWorker),
			JobQueue:  make(chan interface{}, maxBuffer),
			SlefFun: func() {
				fmt.Println("[SelfFun]LoggerDispaatcher start")
			},
		},
	}
	for i := 0; i < maxWorker; i++ {
		dispatcher.Workers[i] = &loggerWorker{}
	}
	return dispatcher
}
