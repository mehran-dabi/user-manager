package main

import (
	"faceit/config"
	"faceit/infrastructure/database"
	"log"
)

func main() {
	//init config
	conf := config.Init()

	// init db
	store, err := database.NewDatabase(conf.Store.Path)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	if err := store.Ping(); err != nil {
		log.Fatalf("failed to get database ping: %s", err)
	}

	if err := store.Migrate("up"); err != nil {
		log.Fatalf("failed to migrate the schemas: %s", err)
	}
}
