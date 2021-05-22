package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/n0v0cain3/go-awsutils/pkg/awsutils"
	"github.com/spf13/viper"
)

func manageUpload(c *fiber.Ctx) error {
	temp, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	file, err := temp.Open()
	if err != nil {
		log.Fatal(err)
	}
	result, err := awsutils.UploadFile(file, temp.Filename)
	if nil != err {
		log.Fatal(err)
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return c.Send([]byte(result.Location))

}

func main() {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(fmt.Errorf("fatal error config file: %s", err))
	}
	app := fiber.New()

	app.Post("/upload", manageUpload)

	if err := app.Listen(":6969"); nil != err {
		log.Fatal(err)
	}
}
