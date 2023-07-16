package commands

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"gman/helpers"
	"strings"
)

func AddAccount() {
	fmt.Println("Enter the GitHub account details:")
	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Token: ")
	var token string
	fmt.Scanln(&token)

	alias := strings.TrimSpace(username)

	credentialName := helpers.GetCredentialName(alias)
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
