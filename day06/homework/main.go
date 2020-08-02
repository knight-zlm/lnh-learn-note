package main

import "github/knight/learn-go/day06/mylogger"

func main() {
	// log := mylogger.NewConsoleLogger("info")
	fp := "."
	fn := "test.log"
	log := mylogger.NewFileLogger("info", fp, fn, 10*1024)
	for i := 0; i < 10; i++ {
		// for {
		log.Debug("这是一条DEBUG日志")
		log.Info("这是一条INFO日志")
		log.Warning("这是一条WARNING日志")
		log.Error("这是一条ERROR日志")
	}
}
