package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     Env.App.Cors,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, apikey",
		AllowCredentials: false,
	}
}
