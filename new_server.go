package gins

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Serve(h http.Handler, host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func() {
		log.Println("serve: " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Println("server exiting")
}
