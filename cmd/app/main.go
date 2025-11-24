package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yporn/doh-ems-api/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ failed to load config: %v", err)
	}	

	log.Printf("🚀 starting %s (env: %s)", cfg.AppName, cfg.Environment)

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	})

	// Route ทดสอบ
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":         cfg.AppName,
			"environment": cfg.Environment,
			"message":     "Config working!",
		})
	})

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("server running at http://%s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
