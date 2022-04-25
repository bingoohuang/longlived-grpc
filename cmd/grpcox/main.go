package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bingoohuang/longlivedgprc/cmd/grpcox/core"
	"github.com/bingoohuang/longlivedgprc/cmd/grpcox/handler"
	"github.com/gorilla/mux"
)

//go:embed index
var indexFS embed.FS

func main() {
	handler.IndexFS = indexFS
	handler.IndexHTML = template.
		Must(template.New("index.html").
			Delims("{[", "]}").
			ParseFS(indexFS, "index/index.html"))

	// logging conf
	//f, err := os.OpenFile("log/grpcox.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer f.Close()
	//log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// start app
	addr := ":6969"
	if value, ok := os.LookupEnv("BIND_ADDR"); ok {
		addr = value
	}
	muxRouter := mux.NewRouter()
	handler.Init(muxRouter)
	var wait time.Duration = time.Second * 15

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      muxRouter,
	}

	fmt.Println("Service started on", addr)
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := removeProtos(); err != nil {
		log.Printf("error while removing protos: %s", err.Error())
	}

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

// removeProtos will remove all uploaded proto file
// this function will be called as the server shutdown gracefully
func removeProtos() error {
	log.Println("removing proto dir from /tmp")
	err := os.RemoveAll(core.BasePath)
	return err
}
