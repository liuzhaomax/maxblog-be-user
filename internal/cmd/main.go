package main

import (
	"context"
	"maxblog-be-template/internal/app"
	"maxblog-be-template/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
