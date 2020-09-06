module github.com/samjegal/fincloud-example

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.11.5-0.20200906111652-5e286818fa7f

require (
	github.com/Azure/go-autorest/autorest v0.11.4 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.2 // indirect
	github.com/samjegal/fincloud-sdk-for-go v1.8.5-0.20200906111302-d5f2fae8c013
	github.com/samjegal/go-fincloud-helpers v0.2.3 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
)
