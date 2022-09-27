package main

import (
	"context"
	"maxblog-be-user/internal/app"
	"maxblog-be-user/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
