package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/samjegal/fincloud-sdk-for-go/services/nas"
)

func main() {
	sender := sender.BuildSender("FINCLOUD")
	client := nas.NewWithKey(
		"9FF16139561D9ED451BA",
		"807D6963AD48E28745DAAD7F846AA3DFE6F27C0B",
	)

	client.Sender = sender

	ctx := context.Background()
	resp, err := nas.Client{BaseClient: client}.GetList(ctx, "", "", "", "", "", "", "", "", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	for _, v := range *resp.GetNasVolumeInstanceListResponse.NasVolumeInstanceList {
		fmt.Println(*v.NasVolumeInstanceNo)
	}
}
