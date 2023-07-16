package commands

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"gman/helpers"
	"os/exec"
	"strings"
)

func SwitchAccount(alias string) {
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

	credentialName := helpers.GetCredentialName(alias)
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
