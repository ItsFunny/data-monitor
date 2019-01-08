package models


// 日志服务器DTO对象,只是一个DTO对象
// 需要传递的信息:
// 1. 日志级别 byte
// 2. 日志信息 string
// FIXME change the name
type Logger struct {
	Level byte
	Msg string
}
