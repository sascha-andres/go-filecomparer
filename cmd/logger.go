package cmd

import (
	"log"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal("Error creating logger")
	}
	sugar = logger.Sugar()
}
