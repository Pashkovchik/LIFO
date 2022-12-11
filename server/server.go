package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"
)

// Server ...
type Server struct {
	Listener net.Listener
}

var server = newServer()
var app = newFiberApp()

func Start() {
	go app.Listener(server.Listener)
	cleanup()
}

func cleanup() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	app.Shutdown()

	os.Exit(1)
}
