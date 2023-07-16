package helpers

import "fmt"

const (
	CredentialTarget = "GitHubAccount"
	ApiURL           = "https://api.github.com"
)

func GetCredentialName(alias string) string {
	return fmt.Sprintf("%s-%s", CredentialTarget, alias)
}
