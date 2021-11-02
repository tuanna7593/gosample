package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/tuanna7593/gosample/app/config"
	"github.com/tuanna7593/gosample/app/external/persistence/mysql"
	"github.com/tuanna7593/gosample/app/external/routes"
)

func main() {
	// load config
	cfg, err := loadConfig("./config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		return
	}

	// init db
	err = mysql.InitDB(cfg.MySQL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return
	}

	// init interrupt signals
	runChan := make(chan os.Signal, 1)

	// set server timeout
	timeOut := cfg.Server.Timeout * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	// Define server
	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: routes.Handler(),
	}
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	// Run the server
	log.Printf("Server is starting on %s\n", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// it's normal case, no need to handler error
			} else {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// block here, wait for the signal to shutdown the server
	interrupt := <-runChan

	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}

func loadConfig(configPath string) (*config.Config, error) {
	// inti config
	cfg := &config.Config{}

	// Open config file
	cfgFile, err := os.Open(configPath)
	if err != nil {
		err = fmt.Errorf("failed to open file %s: %w", configPath, err)
		return nil, err
	}
	defer cfgFile.Close()

	// init new YAML decode
	d := yaml.NewDecoder(cfgFile)

	// decoding from file
	if err := d.Decode(&cfg); err != nil {
		err = fmt.Errorf("failed to parse config file: %w", err)
		return nil, err
	}

	return cfg, nil
}
