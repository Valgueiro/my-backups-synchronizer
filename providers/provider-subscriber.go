package providers

import "fmt"

var SubscribedProviders []Provider = []Provider{
	&GoogleDriveProvider{
		rcloneConfigName: "mateus_personal_drive",
	},
}

func Sync() {
	for _, provider := range SubscribedProviders {
		err := provider.sync()
		if err != nil {
			fmt.Printf("Problem with provider: \n %v", err)
		}
	}
}
