package server

import (
	"github.com/danhkhai07/transfile/app"
)

func main() {
	server := app.Server{}
	server.Run()
}
