package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
)

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
			}),
	}
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(
		func() error {
			err := s.Serve(l)
			if err != nil && err != http.ErrServerClosed {
				log.Printf("falied to close: %+v", err)
				return err
			}
			return nil
		},
	)

	<-ctx.Done()
	err := s.Shutdown(context.Background())
	if err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	return eg.Wait()
}
