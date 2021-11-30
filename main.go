package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	staticPath := os.Getenv("STATIC_PATH")
	serverURL := os.Getenv("SERVER_URL")
	app := fiber.New()
	os.MkdirAll(staticPath, os.ModePerm)
	app.Static("/static", staticPath)
	app.Post("/upload", func(c *fiber.Ctx) error {
		mpf,err := c.MultipartForm()
		if err != nil{
			fmt.Println(err)
			return c.SendStatus(500)
		}
		_ = mpf
		fileHeader := mpf.File["file"][0]
		fileName := fileHeader.Filename // image.png
		fileExtension := fileName[strings.LastIndex(fileName,"."):] // .png
		randomFileName := randSeq(30) + fileExtension
		file, err := fileHeader.Open()
		if err != nil{
			fmt.Println(err)
			return c.SendStatus(500)
		}
		f, err := os.OpenFile(staticPath + randomFileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil{
			fmt.Println(err)
			return c.SendStatus(500)
		}
		defer f.Close()
		io.Copy(f, file)
		return c.SendString(serverURL + randomFileName)
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
