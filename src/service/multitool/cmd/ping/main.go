package main

import (
	"os"

	"github.com/robinbryce/go-layout/service/multitool"
)

func main() {
	multitool.Ping()
	os.Exit(0)
}
