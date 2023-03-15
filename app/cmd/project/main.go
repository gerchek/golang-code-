package main

import (
	"fmt"
	"log"
	"os"
	"project/internal/app"
	"project/internal/setup/constructor"
	"project/pkg/gormclient"
	"project/pkg/logging"

	"github.com/joho/godotenv"
)

func main() {
	// log.Print("logger initializing...")
	logger := logging.Log()

	// project.NewProject(logger)
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	logger.Info("initializing postgres config...")
	gormConfig := gormclient.NewGormConfig(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	logger.Info("connecting to postgres database...")
	client, err := gormclient.NewClient(gormConfig)
	if err != nil {
		logger.Fatal(err)
		fmt.Println(client)
	}

	logger.Info("setting up all repository, service, controller...")

	constructor.SetConstructor(client, logger)

	logger.Info("initializing a new app...")
	app := app.NewApp(logger)

	logger.Fatal(app.Listen(":3000"))

}
