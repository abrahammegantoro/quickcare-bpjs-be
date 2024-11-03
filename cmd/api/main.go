package main

import (
	"context"
	"log"

	"github.com/abrahammegantoro/quickcare-bpjs-be/config"
	"github.com/abrahammegantoro/quickcare-bpjs-be/db"
	"github.com/abrahammegantoro/quickcare-bpjs-be/handlers"
	"github.com/abrahammegantoro/quickcare-bpjs-be/repositories"
	"github.com/abrahammegantoro/quickcare-bpjs-be/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()

	client, err := db.Init(envConfig)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(envConfig.CollectionName)

	app := fiber.New(fiber.Config{
		AppName:      "Quickcare BPJS",
		ServerHeader: "Fiber",
	})

	// repositories
	medicineRepository := repositories.NewMedicineRepository(db)
	referralRepository := repositories.NewReferralRepository(db)

	// services
	medicineService := usecases.NewMedicineUsecase(medicineRepository)
	referralService := usecases.NewReferralUsecase(referralRepository)

	// routes
	server := app.Group("/api/v1").(*fiber.Group)

	// handlers
	handlers.NewMedicineHandler(server, medicineService)
	handlers.NewReferralHandler(server, referralService)

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))

	defer client.Disconnect(context.TODO())
}
