package libraries

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouter() *echo.Echo {
	e := echo.New()
	e.Static("/", "public")
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORS())
	return e
}
