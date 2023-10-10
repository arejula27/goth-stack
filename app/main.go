// main is the entry point
package main

import (
	"html/template"

	"github.com/arejula27/myapp/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	//Load configuration
	config := loadConfig()

	//Create server
	e := echo.New()
	//Load html
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*/*.html")),
	}

	//Css styles
	e.Static("/css", "public/assets")
	e.Renderer = t
	e.GET("/", handlers.Home)

	//Restricted paths
	//r := e.Group("/admin", middlewares.BasicAuth())
	//r.GET("", handlers.PrivateHandler)

	//Start server
	e.Logger.Fatal(e.Start(config.Addres))
}
