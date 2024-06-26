module github.com/beyondstorage/go-storage/services/oss/v3

go 1.16

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.1+incompatible
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/beyondstorage/go-storage/endpoint v1.2.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/google/uuid v1.4.0
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/beyondstorage/go-storage/endpoint => ../../endpoint
	github.com/ragulmathawa/go-storage => ../../
)
