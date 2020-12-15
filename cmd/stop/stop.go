package stop

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	c chan os.Signal
)

func bindSignals(signals ...os.Signal) chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	return c
}

func waitSignal(c <-chan os.Signal) {
	_ = <-c
}

func sendSignal(c chan<- os.Signal, sig os.Signal) {
	c <- sig
}

func Bind() {
	c = bindSignals(syscall.SIGINT, syscall.SIGTERM)
}

func Wait() {
	waitSignal(c)
}

func Stop() {
	sendSignal(c, syscall.SIGINT)
}
