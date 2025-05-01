package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int `json:"code"`
	Message any `json:"msg"`
}

func main() {
	config := GetConfig()
	app := echo.New()

	app.GET("/*", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "Microsoft Store API",
		})
	})

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", config.Port)))
}
