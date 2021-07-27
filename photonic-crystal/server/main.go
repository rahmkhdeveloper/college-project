package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"college-project/gui/photonic-crystal/logic"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/", calculateTransmitivity)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func calculateTransmitivity(c echo.Context) error {
	inputValues := new(logic.Input)

	if err := c.Bind(inputValues); err != nil { // here unmarshal request body into input values
		return c.String(http.StatusInternalServerError, err.Error())
	}
	inputValues.Calculate()
	return c.JSONPretty(http.StatusOK, inputValues, " ")
}
