package Utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const fileMaxSize int = 12
const funcMaxSize int = 20

// Defined Type for the logLevel
type Level int

// Constants Levels available
const (
	LOG Level = iota
	WARN
	ERRO
	FATAL
)

var levelstrings = [...]string{"@", "+@", "!!", "+!!"}

// Traces a log message by level, structured print
// @msg {String}: Main message to display
// @loglvl {Level}: Log level
// @extraMsg {String}: Extra print message
func TraceLog(msg string, loglvl Level, extraMsg string) {
	var printMsg string = levelstrings[loglvl]

	pc, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)

	printMsg += generateTimeStap()
	printMsg += generateCaller(file, runtime.FuncForPC(pc).Name(), line)

	printMsg += fmt.Sprintf(" >>  %s ", msg)
	if extraMsg != "" {
		printMsg += fmt.Sprintf(" { %s }", extraMsg)
	}

	fmt.Println(printMsg)
}

// Generates the log string with the file and line values
// @file {String}: caller file name
func generateCaller(file string, funcname string, line int) string {
	if len(file) > fileMaxSize {
		file = file[:fileMaxSize] + "..."
	}

	splits := strings.Split(funcname, "/")
	funcname = splits[len(splits)-1]
	splits = strings.Split(funcname, ".")
	funcname = splits[len(splits)-1] + "()"

	if len(funcname) > funcMaxSize {
		funcname = funcname[:funcMaxSize] + "..."
	}

	result := fmt.Sprintf(" [ %-15s ] [ %-20s ] [ Line: %-3d ]", file, funcname, line)
	return result
}

// Generates the time stamp for the trace log print
// @Return {String}: Formated time stamp
func generateTimeStap() string {
	cTime := time.Now()

	var result string = fmt.Sprintf(
		" [ %d-%02d-%02d %02d:%02d:%02d:%03d ] ",
		cTime.Year(), cTime.Month(), cTime.Day(),
		cTime.Hour(), cTime.Minute(), cTime.Second(), cTime.UnixNano()*(1.0/int64(time.Millisecond)))

	return result
}
