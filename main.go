package main

import (
	"github.com/labstack/echo/v4"
	"github.com/silverlyjoo/ApiBasic/module"
)

func main() {
	e := echo.New()

	e.GET("/", module.GetData)
	e.GET("/:id", module.GetPath)
	e.GET("/q", module.GetQuery)
	e.POST("/post", module.PostData)

	e.Logger.Fatal(e.Start(":1323"))
}
