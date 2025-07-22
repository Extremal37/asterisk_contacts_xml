package main

import (
	"context"
	"fmt"
	"github.com/Extremal37/asterisk_contacts_xml/internal/app"
	"github.com/Extremal37/asterisk_contacts_xml/internal/config"
	"github.com/Extremal37/asterisk_contacts_xml/internal/logger"
)

func main() {

	// Init Config
	cfg := config.New()

	// Get Config
	err := cfg.ParseConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config: %s", err))
	}

	// Init Logger
	l := logger.New(cfg.Level)
	l.Info("Starting...")

	// Init App
	a := app.New(cfg, l)

	// Launching App
	ctx := context.Background()
	if err = a.Run(ctx); err != nil {
		l.Fatalf("Failed to run app: %s", err)
	}

	l.Info("Work complete!")

}
