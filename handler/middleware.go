package handler

import (
	"fmt"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

func TimeLapsedMiddleware(c *fiber.Ctx) error {
	t := time.Now()

	// Go to next middleware:
	err := c.Next()

	delta := time.Since(t)
	fmt.Println("time lapsed: ", delta)
	return err
}
