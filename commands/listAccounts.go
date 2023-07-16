package commands

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"gman/helpers"
	"strings"
)

func ListAccounts() {
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
		if strings.HasPrefix(cred.TargetName, helpers.CredentialTarget) {
			fmt.Printf("  Username: %s\n", cred.UserName)
			fmt.Println()
		}
	}
}
