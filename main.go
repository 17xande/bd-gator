package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/17xande/bd-gator/internal/config"
	"github.com/17xande/bd-gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("failed to open database connection")
		os.Exit(1)
	}

	dbQueries := database.New(db)

	st := state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

	rawArgs := os.Args
	if len(rawArgs) < 2 {
		fmt.Println("include a command")
		os.Exit(1)
	}

	commandName := rawArgs[1]
	args := rawArgs[2:]
	cmd := command{
		name: commandName,
		args: args,
	}
	err = cmds.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
