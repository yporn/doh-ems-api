package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CustomLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()
		err := c.Next()

		duration := time.Since(start)

		fmt.Printf("[%s] %s %s - Status: %d - Duration: %v\n",
			start.Format("2006-01-02 15:04:05"),
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			duration,
		)

		return err
	}
}
