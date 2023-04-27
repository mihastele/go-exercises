package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)

		select {
		case <-ctx.Done():
			out <- fmt.Sprintf("%v operation cancelled", msg)
			return
		case <-time.After(msDelay * time.Millisecond):
			out <- fmt.Sprintf("%v operation completed", msg)
			return
		}
	}()

	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "web server", 400)
	microService := sampleOperation(ctx, "micro service", 200)
	database := sampleOperation(ctx, "database", 300)

MainLoop:
	for {
		select {
		case msg := <-webServer:
			fmt.Println(msg)
		case msg := <-microService:
			fmt.Println(msg)
			fmt.Println("Cancelling web server operation")
			cancelCtx()
			break MainLoop
		case msg := <-database:
			fmt.Println(msg)
		}
	}

	fmt.Println(<-database)
}
