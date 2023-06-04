package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pouyam79i/Cloud_Computing_HW/main/HW2/step2/code/handler"
)

// This function build and config server for us
func main() {
	fmt.Println("Building Server...")
	e := echo.New()
	e.GET("/", handler.HelloWorld)
	e.POST("/shorten", handler.CallRebrandly)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	fmt.Println("Running Server ...")
	e.Logger.Fatal(e.Start(":8000"))
}
