package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proyecto-integrador/data"
	"proyecto-integrador/routes"
	"syscall"
	"time"
)

func main() {
	// connect to mysql database
	if err := data.ConnectDatabase(); nil != err {
		fmt.Println(err)
		return
	}
	defer data.DB.Close() // program exit close db connection

	// gin router
	router := routes.SetupRouter()

	// service configuration
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Printf("starting server on port %s", srv.Addr)
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shudown the server with a timeout of 5 seconds
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server exiting")
}
