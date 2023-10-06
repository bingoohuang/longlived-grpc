package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bingoohuang/gg/pkg/ctl"
	"github.com/bingoohuang/gg/pkg/fla9"
	_ "github.com/bingoohuang/golog/pkg/autoload"
	"github.com/bingoohuang/golog/pkg/ginlogrus"
	"github.com/gin-gonic/gin"
)

var stoppers Stoppers

func main() {
	pInit := fla9.Bool("init", false, "Initialize a ctl")
	pVersion := fla9.Bool("version", false, "Print version")
	pAddr := fla9.String("addr,a", ":7070", "listen address for the Grpc server")
	pPort := fla9.Int("port,p", 7170, "port for the rest service")
	pMode := fla9.String("mode,m", "both", "client/server/both")
	fla9.Parse()

	ctl.Config{Initing: *pInit, PrintVersion: *pVersion}.ProcessInit()

	gin.SetMode(gin.ReleaseMode)

	gr := gin.New()
	gr.Use(ginlogrus.Logger(nil, true), gin.Recovery())

	switch *pMode {
	case "both", "server":
		addrs := strings.Split(*pAddr, ",")
		for _, addr := range addrs {
			stoppers.Add(startServer(addr))
		}
		gr.GET("/server/:action", serverRestHandle(addrs))
	}

	switch *pMode {
	case "both", "client":
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
