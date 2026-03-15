package config

import (
	"fmt"
	"os"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type StdLogger struct {}

func (ml *StdLogger) Write(s string, args ...any) (error) {
	fmt.Fprint(os.Stdout, time.Now().Format(timeFormat) + " ")
	fmt.Fprintf(os.Stdout, s, args...)
	return nil
}

func (ml *StdLogger) Writeln(s string, args ...any) (error) {
	fmt.Fprint(os.Stdout, time.Now().Format(timeFormat) + " ")
	fmt.Fprintf(os.Stdout, s, args...)
	fmt.Fprintf(os.Stdout, "\n")
	return nil
}

func (ml *StdLogger) Errwrite(s string, args ...any) (error) {
	fmt.Fprint(os.Stderr, time.Now().Format(timeFormat) + " ")
	fmt.Fprintf(os.Stderr, s, args...)
	return nil
}

func (ml *StdLogger) Errwriteln(s string, args ...any) (error) {
	fmt.Fprint(os.Stderr, time.Now().Format(timeFormat) + " ")
	fmt.Fprintf(os.Stderr, s, args...)
	fmt.Fprintf(os.Stderr, "\n")
	return nil
}
