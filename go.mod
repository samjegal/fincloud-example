module github.com/samjegal/fincloud-example

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.11.5-0.20200906111652-5e286818fa7f

require (
	github.com/hashicorp/go-azure-helpers v0.12.0
	github.com/samjegal/fincloud-sdk-for-go v1.9.2-0.20200909093323-beeef8147130
)
