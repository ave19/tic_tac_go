package tictacgo

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace logging
	Trace *log.Logger
	// Info logging
	Info *log.Logger
	// Warning logging
	Warning *log.Logger
	// Error logging
	Error *log.Logger
)

func init() {
	InitLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

//InitLogging stuff
func InitLogging(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	SetTrace(traceHandle)
	SetInfo(infoHandle)
	SetWarning(warningHandle)
	SetError(errorHandle)
}

// SetTrace sets the trace handle
func SetTrace(traceHandle io.Writer) {
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// SetInfo sets the info handle
func SetInfo(infoHandle io.Writer) {
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// SetWarning sets the warning handle
func SetWarning(warningHandle io.Writer) {
	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// SetError sets the error handle
func SetError(errorHandle io.Writer) {
	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
