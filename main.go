package main

import (
	"flag"
	"hg-juke/config"
	"hg-juke/top"
	"log"
	"log/slog"
	"os"
)

func main() {
	logLevel := new(slog.LevelVar)
	logOpts := slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &logOpts))
	slog.SetDefault(logger)

	isDebug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if isDebug != nil && *isDebug {
		logLevel.Set(slog.LevelDebug)
	}

	confExist, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	isInitial := !confExist

	top.Start(isInitial)
}
