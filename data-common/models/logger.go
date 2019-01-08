package models

import "fmt"

const (
	NORMAL = iota
	WARN
	SERIOUS
)

// 日志服务器DTO对象,只是一个DTO对象
// 需要传递的信息:
// 1. 日志级别 byte
// 2. 日志信息 string
// FIXME change the name
type LogDetail struct {
	Level byte
	Msg   string
}

func NewLogDetail(level byte, msg string) *LogDetail {
	return &LogDetail{Level: level, Msg: msg}
}

type LogModule struct {
	logDetail  LogDetail
	moduleName string //
}

func (l LogDetail) String() string {
	return fmt.Sprintf("level=%d,msg=%s", l.Level, l.Msg)
}
