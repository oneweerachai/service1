package main

import "fmt"
import  log "github.com/oneweerachai/shared-logger"

func main() {
	logger,err := log.NewLogger()

	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	fmt.Println("Hello")

	logger.Info("Hello, shared logger")
}
