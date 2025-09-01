package main

import (
	"os"

	"github.com/dstm45/vacances/pkg/api"
)

func main() {
	config := api.Config{
		Addr: os.Getenv("Addr"),
		Dsn:  "postgres://postgres:yourpassword@localhost:5432/vacances",
	}
	server := api.NewServer(config)
	server.Start()
}
