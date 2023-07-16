package commands

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"gman/helpers"
)

func RemoveAccount(alias string) {
	// prompt for confirmation
	fmt.Printf("Are you sure you want to delete '%s' GitHub account? (y/n) ", alias)
	var confirmation string
	fmt.Scanln(&confirmation)
	if confirmation != "y" {
		fmt.Printf("%s GitHub account not removed.\n", alias)
		return
	}
	credentialName := helpers.GetCredentialName(alias)
	cred, err := wincred.GetGenericCredential(credentialName)
	if err != nil {
		fmt.Println("Failed to retrieve GitHub account details:", err)
		return
	}
	cred.Delete()

	fmt.Printf("Deleted '%s' GitHub account.\n", alias)
}
