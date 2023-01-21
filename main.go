package main

import (
	"context"
	"log"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("falied to terminate server: %v", err)
	}
}
