module github.com/beyondstorage/go-storage/services/kodo/v3

go 1.16

require (
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/beyondstorage/go-storage/endpoint v1.2.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/google/uuid v1.4.0
	github.com/qiniu/go-sdk/v7 v7.20.0
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/beyondstorage/go-storage/endpoint => ../../endpoint
	github.com/ragulmathawa/go-storage => ../../
)
