package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	logger := New(os.Stdout, "new: ", 0)

	logger.FatalfIf(len(os.Args) < 2, "Subcommand name argument expected")
	subcommand := os.Args[1]
	switch subcommand {
	case "scala", "scala-compiler":
		re := regexp.MustCompile(`[\d]+\.[\d]+\.[\d]+`)
		basename := filepath.Base(os.Getenv("SCALA_HOME"))
		fmt.Printf("%s\n", re.FindString(basename))
	}
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
