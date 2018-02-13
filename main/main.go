package main

import (
	"context"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/ext-dns"
)

func main() {
	ctx := context.Background()
	funcServer := server.NewFromEnv(ctx)
	funcServer.AddExtension(&dns.Dns{})
	funcServer.Start(ctx)
}
