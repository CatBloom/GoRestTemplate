package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	log.Println("init")
	e = echo.New()
}

func main() {
	e.Start(":8080")
}
