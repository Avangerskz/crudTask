package main

import (
	"context"
	"embed"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"taskRestAPI/configs"
	"taskRestAPI/internal/app"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, 1)
	go app.StartHTTPServer(ctx, cfg, errCh)

	log.Fatalf("Terminated: %s", <-errCh)
}
