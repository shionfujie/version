package main

import (
	"os/exec"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var semverRe = regexp.MustCompile(`[\d]+\.[\d]+\.[\d]+`)

func main() {
	logger := New(os.Stdout, "version: ", 0)

	logger.FatalfIf(len(os.Args) < 2, "Subcommand name argument expected")
	subcommand := os.Args[1]
	switch subcommand {
	case "scala", "scala-compiler":
		basename := filepath.Base(os.Getenv("SCALA_HOME"))
		fmt.Println(semverRe.FindString(basename))
	case "go":
		logger.SetPrefix("version go: ")
		if _, err := exec.LookPath("go"); err != nil {
			logger.Fatalln("'go' executable expected to be available by the PATH environment variable")
		}
		o, err := exec.Command("go", "version").Output()
		if err != nil {
			logger.Fatalln("Failed to run 'go version'")
		}
		fmt.Printf("%s\n", semverRe.Find(o))
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
