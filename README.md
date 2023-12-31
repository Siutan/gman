
# GitHub Account Manager (gman)

gman is a command-line tool for managing multiple GitHub accounts on Windows. It allows you to easily switch between different GitHub accounts, add new accounts, and remove existing accounts. The tool securely stores the account details in the Windows Credential Manager.

## Prerequisites

- Windows operating system
- Git command-line tool installed and accessible in the system's PATH
- Go programming language (optional, if you want to build the script from source)

## Installation

1. Download the latest release of gman from the [Releases](https://github.com/Siutan/gman/releases) page.
2. Extract the downloaded archive to a location of your choice.
3. Add the location of the extracted archive to the system's PATH environment variable if you want to be able to run the script from any location in the command-line.

Alternatively, you can build the script from source using the Go programming language:

```
go build -o gman gman.go
```

## Usage

### Adding an Account

To add a new GitHub account:

```
gman add
```

Follow the prompts and enter the username, and personal access token for the GitHub account. The username is used to identify the account. Find out how to get your access token [Here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens).

### Listing Accounts

To list all the GitHub accounts stored by gman:

```
gman list
```
This will return all the usernames of the stored accounts.

### Switching an Account

To switch to a specific GitHub account:

```
gman switch <username>
```

Replace `<username>` with the username assigned to the account you want to switch to.

### Removing an Account

To remove an existing GitHub account:

```
gman remove <username>
```

Replace `<username>` with the username assigned to the account you want to remove. You will be asked for confirmation before the account is deleted.

### List Available Commands

To see the available commands and their usage:

```
gman

```

### Check version

To see the current verison of gman run:
```
gman version

```

## Security

gman securely stores the GitHub account details (username and personal access token) in the Windows Credential Manager.
Encryption will be added in a future release.

## To-Do

- [ ] Add encryption for the stored account details
- [ ] Allow backing up and restoring the stored account details
- [ ] Add support for aliases
- [ ] Add edit command to edit the stored account details
- [x] Add list command to list all the stored accounts
- [ ] Add repo config command to see which account is being used for a repository
## License

This project is licensed under the [MIT License](LICENSE).
