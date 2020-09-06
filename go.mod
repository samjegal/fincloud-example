module github.com/samjegal/fincloud-example

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.11.5-0.20200906111652-5e286818fa7f

require (
	github.com/hashicorp/go-azure-helpers v0.12.0
	github.com/samjegal/fincloud-sdk-for-go v1.8.5-0.20200906111302-d5f2fae8c013
	github.com/samjegal/go-fincloud-helpers v0.2.4 // indirect
)
