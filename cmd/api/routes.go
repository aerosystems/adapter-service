package main

import "github.com/labstack/echo/v4"

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()
	return e
}
