package config

import (
	"fmt"
)

type Logger interface {
	Write(s string)
	Writeln(s string)
}

type StdoutLogger struct {}

func (ml *StdoutLogger) Write(s string) (error) {
	fmt.Printf(s)
	return nil
}

func (ml *StdoutLogger) Writeln(s string) (error) {
	return nil
}
