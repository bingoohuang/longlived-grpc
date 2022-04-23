package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bingoohuang/gg/pkg/ctl"
	"github.com/bingoohuang/gg/pkg/fla9"
	"github.com/bingoohuang/golog"
	"github.com/bingoohuang/golog/pkg/ginlogrus"
	"github.com/gin-gonic/gin"
)

var stoppers Stoppers

func main() {
	pVerbose := fla9.Bool("verbose,v", false, "Initialize a ctl")
	pInit := fla9.Bool("init", false, "Initialize a ctl")
	pVersion := fla9.Bool("version", false, "Print version")
	pAddr := fla9.String("addr,a", ":7070", "listen address for the Grpc server")
	pPort := fla9.Int("port,p", 7170, "port for the rest service")
	pMode := fla9.String("mode,m", "both", "client/server/both")
	fla9.Parse()

	ctl.Config{Initing: *pInit, PrintVersion: *pVersion}.ProcessInit()

	gin.SetMode(gin.ReleaseMode)
	var fns []golog.SetupOptionFn
	if *pVerbose {
		fns = append(fns, golog.Spec("stdout"))
	}
	golog.Setup(fns...)

	gr := gin.New()
	gr.Use(ginlogrus.Logger(nil, true), gin.Recovery())

	if *pMode == "both" || *pMode == "server" {
		addrs := strings.Split(*pAddr, ",")
		for _, addr := range addrs {
			stoppers.Add(startServer(addr))
		}
		gr.GET("/server/:action", serverRestHandle(addrs))
	}
	if *pMode == "both" || *pMode == "client" {
		gr.GET("/client/:action", clientRestHandle(*pAddr))
	}

	go func() {
		addr := fmt.Sprintf(":%d", *pPort)
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
