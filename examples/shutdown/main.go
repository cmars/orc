package main

import (
	"github.com/sycamoreone/orc/control"
	"log"
)

func main() {
	c, err := control.Dial(":9051")
	if err != nil {
		log.Print(err)
	}
	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("connected; shuting down again")

	err = c.Signal(control.SignalShutdown)
	if err != nil {
		log.Print(err)
		return
	}
	reply, err := c.Receive()
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(reply)
}
