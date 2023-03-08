package Utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
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

	printMsg += fmt.Sprintf(" [ %02d:%02d:%02d:%03dms ]", 1, 1, 10, 300)

	pc, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	printMsg += generate_caller(file, runtime.FuncForPC(pc).Name(), line)

	printMsg += fmt.Sprintf(" >>  %s ", msg)
	if extraMsg != "" {
		printMsg += fmt.Sprintf(" { %s }", extraMsg)
	}

	fmt.Println(printMsg)
}

// Generates the log string with the file and line values
// @file {String}: caller file name
func generate_caller(file string, funcname string, line int) string {
	if len(file) > fileMaxSize {
		file = file[:fileMaxSize] + "..."
	}

	splits := strings.Split(funcname, "/")
	funcname = splits[len(splits)-1]
	splits = strings.Split(funcname, ".")
	funcname = splits[len(splits)-1] + "( )"

	if len(funcname) > funcMaxSize {
		funcname = funcname[:funcMaxSize] + "..."
	}

	result := fmt.Sprintf(" [ %-15s ] [ %-20s ] [ Line: %-3d ]", file, funcname, line)
	return result
}
