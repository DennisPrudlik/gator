package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DennisPrudlik/gator/internal/config"
	"github.com/DennisPrudlik/gator/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db      *database.Queries
	cfg_ptr *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	command_handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing arguments")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		os.Exit(1)
	}
	err = s.cfg_ptr.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %s logged in\n", s.cfg_ptr.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing arguments")
	}
	// Registration logic here
	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		Name:      cmd.args[0],
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		ID:        uuid.New(),
	})
	if err != nil {
		os.Exit(1)
	}
	err = s.cfg_ptr.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User registered: %v\n", usr)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.command_handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(*state, command) error) {
	if c.command_handlers == nil {
		c.command_handlers = make(map[string]func(*state, command) error)
	}
	c.command_handlers[name] = handler
}
