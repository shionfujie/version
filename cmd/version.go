package main

import (
	"io"
	"log"
	"regexp"
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	re := regexp.MustCompile(`[\d]+\.[\d]+\.[\d]+`)
	basename := filepath.Base(os.Getenv("SCALA_HOME"))
	fmt.Printf("%s\n", re.FindString(basename))
}

type sLogger struct {
	*log.Logger
	O io.Writer
}

func New(out io.Writer, prefix string, flag int) *sLogger {
	return &sLogger{log.New(out, prefix, flag), out}
}

func (l *sLogger) FatalfIf(b bool, format string, a ...interface{}) {
	if b {
		l.Fatalf(format+"\n", a...)
	}
}

func (l *sLogger) FatalfIfError(err error, format string, a ...interface{}) {
	l.FatalfIf(err != nil, format, a...)
}