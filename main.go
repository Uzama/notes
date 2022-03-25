package main

import (
	"context"
	"notes/bootstrap"
)

func main() {
	ctx := context.Background()

	bootstrap.Start(ctx)
}
