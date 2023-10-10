// main is the entry point
package main

import (
	"github.com/arejula27/myapp/handlers"
	"github.com/arejula27/myapp/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {

	//Load configuration
	config := loadConfig()
	e := echo.New()
	e.GET("/", handlers.PublicHandler)

	//Restricted paths
	r := e.Group("/admin", middlewares.BasicAuth())
	r.GET("", handlers.PrivateHandler)
	e.Logger.Fatal(e.Start(config.Addres))
}
