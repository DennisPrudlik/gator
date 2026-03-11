package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DennisPrudlik/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := state{cfg_ptr: &cfg}

	c := commands{}
	c.register("login", handlerLogin)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No command provided")
		os.Exit(1)
	}
	cmd := command{name: args[0], args: args[1:]}

	err = c.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
