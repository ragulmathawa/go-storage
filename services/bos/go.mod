module github.com/beyondstorage/go-storage/services/bos/v2

go 1.16

require (
	github.com/baidubce/bce-sdk-go v0.9.176
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/beyondstorage/go-storage/endpoint v1.2.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/google/uuid v1.6.0
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/beyondstorage/go-storage/endpoint => ../../endpoint
	github.com/ragulmathawa/go-storage => ../../
)
