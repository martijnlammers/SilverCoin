package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Startup validation.

	path_config := flag.String("config", "", "specify configuration path")
	flag.Parse()
	if _, err := os.Stat(*path_config); err != nil {
		fmt.Println("invalid path to config")
		return
	}

	// fmt.Println("file exists")
	// test := []byte("hello")
	// fmt.Println(scm.Hash(test))

	// Server and API section

	app := fiber.New()
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("It works!")
	})
	app.Listen(":8000")
}
