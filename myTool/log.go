package myTool

import (
	"io"
	"log"
)

//Ldate         = 1 << iota     // 形如 2009/01/23 的日期
//Ltime                         // 形如 01:23:23   的时间
//Lmicroseconds                 // 形如 01:23:23.123123   的时间
//Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
//Lshortfile                    // 文件名和行号: d.go:23
//LstdFlags     = Ldate | Ltime // 日期和时间

func New(logFile io.Writer, prefix string, logType int) *log.Logger {
	var debugLog *log.Logger
	switch logType {
	case 1:
		debugLog = log.New(logFile, prefix, log.Llongfile)
		break
	case 2:
		debugLog = log.New(logFile, prefix, log.Ltime)
		break
	case 3:
		debugLog = log.New(logFile, prefix, log.Lmicroseconds)
		break
	case 4:
		debugLog = log.New(logFile, prefix, log.Ldate)
		break
	case 5:
		debugLog = log.New(logFile, prefix, log.Lshortfile)
		break
	case 6:
		debugLog = log.New(logFile, prefix, log.LstdFlags)
		break
	}
	return debugLog
}
