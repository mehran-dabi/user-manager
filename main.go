package main

import (
	"context"
	"faceit/config"
	"faceit/domain/controller"
	"faceit/domain/service"
	"faceit/domain/user/repository"
	"faceit/infrastructure/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//init config
	conf := config.Init()

	// init db
	store, err := database.NewDatabase(
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
		conf.Database.Driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	if err := store.Ping(); err != nil {
		log.Fatalf("failed to get database ping: %s", err)
	}

	if err := store.Migrate("up"); err != nil {
		log.Fatalf("failed to migrate the schemas: %s", err)
	}

	usersRepo := repository.NewUserRepository(store.DB)
	usersService := service.NewUserService(usersRepo)
	usersController := controller.NewUserController(usersService)

	server := usersController.Run(conf.Service.Port)

	waitForOsSignal()
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func waitForOsSignal() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-osSignal
}
