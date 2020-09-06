package main

import (
	"context"
	"fmt"

	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
)

func main() {
	// sender := sender.BuildSender("FINCLOUD")
	client := vpc.NewClient()

	client.AccessKey = "9FF16139561D9ED451BA"
	client.Secretkey = "807D6963AD48E28745DAAD7F846AA3DFE6F27C0B"
	// client.Sender = sender

	ctx := context.Background()
	resp, err := client.GetList(ctx, "", "", "", "")
	if err != nil {
		panic("send request faild...")
	}

	fmt.Println(resp.StatusCode)
	for _, v := range *resp.GetVpcListResponse.VpcList {
		fmt.Println(*v.VpcName)
	}
}
