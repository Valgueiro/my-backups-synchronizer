package main

import (
	"fmt"
	"my-backups-synchronizer/providers"
)

func main() {
	fmt.Println("Hello! ------")

	fmt.Println("Syncing Providers...")
	providers.Sync()
}
