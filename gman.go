package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/danieljoos/wincred"
)

const (
	credentialTarget = "GitHubAccount"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: go gman <command> [options]")
		fmt.Println("Available commands:")
		fmt.Println("  switch <alias> - Switch to the specified GitHub account")
		fmt.Println("  add            - Add a new GitHub account")
		fmt.Println("  remove <alias> - Remove an existing GitHub account")
		return
	}

	command := os.Args[1]

	switch command {
	case "switch":
		if len(os.Args) < 2 {
			fmt.Println("Usage: go gman switch <alias>")
			return
		}
		alias := os.Args[2]
		switchAccount(alias)
	case "add":
		addAccount()
	case "remove":
		if len(os.Args) < 2 {
			fmt.Println("Usage: go gman remove <alias>")
			return
		}
		alias := os.Args[2]
		removeAccount(alias)
	default:
		fmt.Println("Invalid command.")
	}
}

func switchAccount(alias string) {
	credentialName := getCredentialName(alias)
	cred, err := wincred.GetGenericCredential(credentialName)
	if err != nil {
		fmt.Println("Failed to retrieve GitHub account details:", err)
		return
	}

	if cred == nil {
		fmt.Printf("GitHub account '%s' not found.\n", alias)
		return
	}

	username := cred.UserName
	token := cred.CredentialBlob

	cmd := exec.Command("git", "config", "--global", "user.name", username)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to switch GitHub account:", err)
		return
	}

	cmd = exec.Command("git", "config", "--global", "user.password", string(token))
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to switch GitHub account:", err)
		return
	}

	fmt.Printf("Switched to '%s' GitHub account.\n", alias)
}

func addAccount() {
	fmt.Println("Enter the GitHub account details:")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Token: ")
	var token string
	fmt.Scanln(&token)

	alias := strings.TrimSpace(username)

	credentialName := getCredentialName(alias)
	cred := wincred.NewGenericCredential(credentialName)
	cred.UserName = username
	cred.CredentialBlob = []byte(token)
	err := cred.Write()
	if err != nil {
		fmt.Println("Failed to add GitHub account:", err)
		return
	}

	fmt.Printf("Added '%s' GitHub account.\n", alias)
}

func removeAccount(alias string) {
	credentialName := getCredentialName(alias)
	cred, err := wincred.GetGenericCredential(credentialName)
	if err != nil {
		fmt.Println("Failed to retrieve GitHub account details:", err)
		return
	}
	cred.Delete()

	fmt.Printf("Deleted '%s' GitHub account.\n", alias)
}

func getCredentialName(alias string) string {
	return fmt.Sprintf("%s-%s", credentialTarget, alias)
}
