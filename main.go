package main

import (
	"github.com/mvrcoag/service-monitor/cmd"
	"github.com/mvrcoag/service-monitor/storage"
)

func main() {
	storage.InitStorage()
	cmd.Execute()
}
