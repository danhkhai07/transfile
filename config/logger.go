package config

import (
	"fmt"
	"os"
)

type Logger interface {
	Write(s string, args ...any) (error)
	Writeln(s string, args ...any) (error)
	Errwrite(s string, args ...any) (error)
	Errwriteln(s string, args ...any) (error)
}

type StdLogger struct {}

func (ml *StdLogger) Write(s string, args ...any) (error) {
	fmt.Fprintf(os.Stdout, s, args...)
	return nil
}

func (ml *StdLogger) Writeln(s string, args ...any) (error) {
	fmt.Fprintf(os.Stdout, s, args...)
	fmt.Fprintf(os.Stdout, "\n")
	return nil
}

func (ml *StdLogger) Errwrite(s string, args ...any) (error) {
	fmt.Fprintf(os.Stderr, s, args...)
	return nil
}

func (ml *StdLogger) Errwriteln(s string, args ...any) (error) {
	fmt.Fprintf(os.Stderr, s, args...)
	fmt.Fprintf(os.Stderr, "\n")
	return nil
}
