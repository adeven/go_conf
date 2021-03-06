package go_conf

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ExitHandler interface {
	OnExit()
}

type StandardHandler struct {
}

func (self *StandardHandler) OnExit() {
}

func signalCatcher() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	for signal := range ch {
		if signal == syscall.SIGHUP {
			log.Println("received SIGHUP exiting...")
			exitHandler.OnExit()
			os.Exit(0)
		}
	}
}

func startSignalCatcher() {
	//react to sighup
	go signalCatcher()
}

func SetExitHandler(handler ExitHandler) {
	exitHandler = handler
}
