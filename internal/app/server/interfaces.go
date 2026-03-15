package app

import (

)

type Logger interface {
	Write(s string, args ...any) (error)
	Writeln(s string, args ...any) (error)
	Errwrite(s string, args ...any) (error)
	Errwriteln(s string, args ...any) (error)
}

