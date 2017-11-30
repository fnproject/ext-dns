package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	"github.com/treeder/fn-ext-cloudflare"
)

func main() {
	ctx := context.Background()

	funcServer := server.NewFromEnv(ctx)

	funcServer.AddRootMiddleware(&dns.Middleware{})

	funcServer.Start(ctx)
}
