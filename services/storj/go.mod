module github.com/beyondstorage/go-storage/services/storj

go 1.16

require (
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/google/uuid v1.4.0
	storj.io/uplink v1.12.2
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/ragulmathawa/go-storage => ../../
)
