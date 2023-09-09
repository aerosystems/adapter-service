package main

import (
	"fmt"
)

const webPort = "80"

func main() {
	app := NewApp()

	e := app.NewRouter()

	err := e.Start(fmt.Sprintf(":%s", webPort))
	if err != nil {
		panic(err)
	}
}
