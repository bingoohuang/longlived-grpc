package main

import (
	"fmt"
	"github.com/bingoohuang/gg/pkg/ctl"
	"github.com/bingoohuang/gg/pkg/fla9"
	"github.com/bingoohuang/golog"
	"github.com/bingoohuang/golog/pkg/ginlogrus"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
)

var stoppers Stoppers

func main() {
	pInit := fla9.Bool("init", false, "Initialize a ctl")
	pVersion := fla9.Bool("version,v", false, "Print version")
	pAddr := fla9.String("addr,a", ":7070", "listen address for the Grpc server")
	pMode := fla9.String("mode,m", "both", "client/server/both")
	fla9.Parse()

	ctl.Config{Initing: *pInit, PrintVersion: *pVersion}.ProcessInit()

	gin.SetMode(gin.ReleaseMode)
	golog.Setup()
	gr := gin.New()
	gr.Use(ginlogrus.Logger(nil, true), gin.Recovery())

	host, sport, err := net.SplitHostPort(*pAddr)
	if err != nil {
		log.Fatalf("parse host and port from argument addr, failed: %v", err)
	}

	port, _ := strconv.Atoi(sport)
	port += 10
	if *pMode == "both" || *pMode == "server" {
		stoppers.Add(startServer(*pAddr))
		gr.GET("/server/:action", serverRestHandle(*pAddr))
	}
	if *pMode == "both" || *pMode == "client" {
		gr.GET("/client/:action", clientRestHandle(*pAddr))
	}

	if *pMode == "client" {
		port++
	}

	go func() {
		addr := fmt.Sprintf("%s:%d", host, port)
		log.Printf("ListenAndServe rest server at %s", addr)
		if err := gr.Run(addr); err != nil {
			log.Printf("E! ListenAndServe rest server failed: %v", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	log.Printf("signal %v received", <-c)
	stoppers.Stop()
	log.Print("exit")
}
