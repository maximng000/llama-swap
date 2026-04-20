package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mostlygeek/llama-swap/proxy"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configFile  = flag.String("config", "config.yaml", "path to configuration file")
		listenAddr  = flag.String("listen", ":11434", "address to listen on") // use ollama's default port
		showVersion = flag.Bool("version", false, "print version information and exit")
		logLevel    = flag.String("log-level", "info", "log level (debug, info, warn, error)")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("llama-swap version %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	if *configFile == "" {
		log.Fatal("config file is required")
	}

	cfg, err := proxy.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Printf("llama-swap %s starting on %s", version, *listenAddr)
	log.Printf("loaded config from %s", *configFile)

	if *logLevel == "debug" {
		log.Printf("debug logging enabled")
	}

	server, err := proxy.NewServer(cfg, *listenAddr)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
