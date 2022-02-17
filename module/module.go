package module

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetData(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}

func GetPath(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetQuery(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	return c.String(http.StatusOK, "name : "+name+", age : "+age)
}

func PostData(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	return c.String(http.StatusOK, "name : "+name+", age : "+age)
}
