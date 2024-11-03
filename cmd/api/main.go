package main

import (
	"log"

	"github.com/abrahammegantoro/quickcare-bpjs-be/config"
	"github.com/abrahammegantoro/quickcare-bpjs-be/db"
	"github.com/abrahammegantoro/quickcare-bpjs-be/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()

	client, err := db.Init(envConfig)
	if err != nil {
		log.Fatal("Error initializing database: %e", err)
	}

	db := client.Database(envConfig.CollectionName)

	app := fiber.New(fiber.Config{
		AppName:      "Quickcare BPJS",
		ServerHeader: "Fiber",
	})

	// repositories
	medicineRepository := repositories.NewMedicineRepository(db)

}
