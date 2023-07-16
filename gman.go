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
		fmt.Println("  switch <username> - Switch to the specified GitHub account")
		fmt.Println("  add            - Add a new GitHub account")
		fmt.Println("  remove <username> - Remove an existing GitHub account")
		return
	}

	command := os.Args[1]

	switch command {
	case "switch":
		if len(os.Args) < 2 {
			fmt.Println("Usage: gman switch <username>")
			return
		}
		alias := os.Args[2]
		switchAccount(alias)
	case "add":
		addAccount()
	case "remove":
		if len(os.Args) < 2 {
			fmt.Println("Usage: gman remove <username>")
			return
		}
		alias := os.Args[2]
		removeAccount(alias)
	case "list":
		listAccounts()
	default:
		fmt.Println("Invalid command.")
	}
}

func switchAccount(alias string) {
	// if current account is the same as the one we're switching to, do nothing
	checkCmd := exec.Command("git", "config", "--global", "user.name")
	currentUsername, err := checkCmd.Output()
	if err != nil {
		fmt.Println("Failed to switch GitHub account:", err)
		return
	}
	currentUsername = []byte(strings.TrimSpace(string(currentUsername)))
	if string(currentUsername) == alias {
		fmt.Printf("Already using '%s' GitHub account.\n", alias)
		return
	}

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
	// prompt for confirmation
	fmt.Printf("Are you sure you want to delete '%s' GitHub account? (y/n) ", alias)
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Printf("%s GitHub account not removed.\n", alias)
		return
	}
	credentialName := getCredentialName(alias)
	cred, err := wincred.GetGenericCredential(credentialName)
	if err != nil {
		fmt.Println("Failed to retrieve GitHub account details:", err)
		return
	}
	cred.Delete()

	fmt.Printf("Deleted '%s' GitHub account.\n", alias)
}

func listAccounts() {
	creds, err := wincred.List()
	if err != nil {
		fmt.Println("Failed to retrieve GitHub accounts:", err)
		return
	}

	if len(creds) == 0 {
		fmt.Println("No GitHub accounts found.")
		return
	}

	fmt.Println("GitHub Accounts:")
	for _, cred := range creds {
		if strings.HasPrefix(cred.TargetName, credentialTarget) {
			fmt.Printf("  Username: %s\n", cred.UserName)
			fmt.Println()
		}
	}
}

func getCredentialName(alias string) string {
	return fmt.Sprintf("%s-%s", credentialTarget, alias)
}
