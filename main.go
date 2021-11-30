package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	staticPath := os.Getenv("STATIC_PATH")
	app := fiber.New()

	app.Static("/static", staticPath)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Post("/upload", func(c *fiber.Ctx) error {
		if len(c.Body()) > 0 {
			randomFileName := randSeq(30) + ".png"
			ioutil.WriteFile("static/"+randomFileName, c.Body(), 0644)
			return c.SendString("http://localhost:3000/img/" + randomFileName)
		}
		return c.SendStatus(500)
	})
	app.Listen(":3000")
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
