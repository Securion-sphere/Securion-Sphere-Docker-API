package main

import (
	"fmt"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/bootstrap"
)

func main() {
	app, cfg := bootstrap.Bootstrap()
	if app == nil {
		return
	}
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.AppPort)))
}
