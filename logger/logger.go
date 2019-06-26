package logger

import (
	"fmt"

	"github.com/nektro/go-util/util"
)

//

type LogLevel int

const (
	LevelALL LogLevel = iota
	LeveLTRACE
	LevelDEBUG
	LevelINFO
	LevelWARN
	LevelERROR
	LevelFATAL
)

//

type Logger struct {
	Level   LogLevel
	Enabled bool
}

//

func New() Logger {
	return Logger{LevelINFO, true}
}

//

func (l *Logger) Log(lev LogLevel, message ...interface{}) {
	if l.Enabled && lev >= l.Level {
		fmt.Print("[" + util.GetIsoDateTime() + "] ")
		fmt.Println(message...)
	}
}
