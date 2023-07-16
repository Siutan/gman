package commands

import "fmt"

var CurrVer = "dev"

func Version() {
	fmt.Printf("gman %s\n", CurrVer)
}
