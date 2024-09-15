package main

import (
	"fmt"

	log "github.com/oneweerachai/shared-logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger,err := log.NewLoggerWithConfig(zapcore.DebugLevel)

	defer logger.Sync()

	if err != nil {
		panic(err)
	}
	fmt.Println("Hello")

	logger.Info("Hello, shared logger")
	logger.Debug("h debug")
}
