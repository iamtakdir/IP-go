package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtakdir/pixels-go/models"
)



func main() {
	app := fiber.New()

	app.Get("/ip::ip", GetIp)

	app.Listen(":" + os.Getenv("PORT"))
}

func GetIp(c *fiber.Ctx) error {

	baseIp := c.Params("ip")

	url := "http://api.ipstack.com/" + baseIp + "?access_key=" + Token + "&format=1"
	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	respData, err := ioutil.ReadAll(response.Body)
	var fire_error models.Error

	if err != nil {
		json.Unmarshal(respData, &fire_error)
	}

	var respObj models.IpTable

	json.Unmarshal(respData, &respObj)

	return c.JSON(respObj)
}
