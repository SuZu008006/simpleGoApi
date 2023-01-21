package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
			}),
	}
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(
		func() error {
			err := s.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Printf("failed to close: %+v", err)
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
