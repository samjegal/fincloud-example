package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/samjegal/fincloud-sdk-for-go/services/nas"
)

func main() {
	sender := sender.BuildSender("FINCLOUD")
	client := nas.NewClient()

	client.AccessKey = "9FF16139561D9ED451BA"
	client.Secretkey = "807D6963AD48E28745DAAD7F846AA3DFE6F27C0B"

	client.Sender = sender

	ctx := context.Background()
	resp, _ := client.GetList(ctx, "", "", "", "", "", "", "", "", "", "")

	fmt.Println(resp.StatusCode)

	for _, v := range *resp.GetNasVolumeInstanceListResponse.NasVolumeInstanceList {
		fmt.Println(*v.NasVolumeInstanceNo)
	}
}
