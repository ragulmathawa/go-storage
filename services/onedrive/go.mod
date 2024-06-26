module github.com/beyondstorage/go-storage/services/onedrive

go 1.16

require (
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/goh-chunlin/go-onedrive v1.1.1
	github.com/google/uuid v1.4.0
	golang.org/x/oauth2 v0.15.0
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/ragulmathawa/go-storage => ../../
)
