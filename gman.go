package main

import (
	"fmt"
	"gman/commands"
	"os"
)

// Todo:
// - Add support for Linux and Mac
// - Add support for listing accounts
// - Add support for init command

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: go gman <command> [options]")
		fmt.Println("Available commands:")
		fmt.Println("  switch <username> - Switch to the specified GitHub account")
		fmt.Println("  add            - Add a new GitHub account")
		fmt.Println("  remove <username> - Remove an existing GitHub account")
		return
	}

	command := os.Args[1]

	switch command {
	case "version":
		commands.Version()
	case "switch":
		if len(os.Args) < 2 {
			fmt.Println("Usage: gman switch <username>")
			return
		}
		alias := os.Args[2]
		commands.SwitchAccount(alias)
	case "add":
		commands.AddAccount()
	case "remove":
		if len(os.Args) < 2 {
			fmt.Println("Usage: gman remove <username>")
			return
		}
		alias := os.Args[2]
		commands.RemoveAccount(alias)
	case "list":
		commands.ListAccounts()
	default:
		fmt.Println("Invalid command.")
	}
}
