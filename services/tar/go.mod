module github.com/beyondstorage/go-storage/services/tar

go 1.16

require (
	github.com/beyondstorage/go-storage/endpoint v1.2.0
	github.com/ragulmathawa/go-storage v5.0.0
	github.com/stretchr/testify v1.9.0
)

replace (
	github.com/beyondstorage/go-storage/endpoint => ../../endpoint
	github.com/ragulmathawa/go-storage => ../../
)
