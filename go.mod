module github.com/samjegal/fincloud-example

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.0.0-20200908011700-76a89b3e2bac

require (
	github.com/hashicorp/go-azure-helpers v0.12.0
	github.com/samjegal/fincloud-sdk-for-go v1.9.1
)
