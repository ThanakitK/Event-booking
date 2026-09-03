package main

import (
	"event-booking/config"
	"event-booking/core/handlers"
	"event-booking/core/repositories"
	"event-booking/core/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	config.NewAppInitEnvironment()
}
func main() {
	db := config.NewAppDatabase()

	masterDB := db.Client().Database(config.Env.DB.Name)
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))

	// Repositories
	eventRepo := repositories.NewEventRepository(masterDB, "events")
	userRepo := repositories.NewUserRepository(masterDB, "users")

	// Services
	eventService := services.NewEventService(eventRepo)
	userService := services.NewUserService(userRepo)

	// Handlers
	eventHand := handlers.NewEventHandler(eventService)
	userHand := handlers.NewUserHandler(userService)

	api := app.Group("/api")

	// events
	api.Post("/event", eventHand.CreateEvent)
	api.Put("/event/:id", eventHand.UpdateEvent)
	api.Delete("/event/:id", eventHand.DeleteEvent)

	// users
	api.Post("/user", userHand.CreateUser)

	log.Fatal(app.Listen(":" + config.Env.App.Port))
}
