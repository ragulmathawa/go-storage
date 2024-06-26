module github.com/beyondstorage/go-storage/services/cos/v3

go 1.16

require (
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/google/uuid v1.4.0
	github.com/tencentyun/cos-go-sdk-v5 v0.7.45
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/ragulmathawa/go-storage => ../../
)
