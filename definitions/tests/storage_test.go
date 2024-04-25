package main

import (
	"context"
	"io"

	"github.com/ragulmathawa/go-storage/types"
)

type Storage struct {
	defaultPairs types.DefaultStoragePairs
	features     types.StorageFeatures

	Pairs []types.Pair

	types.UnimplementedStorager
}

func (s *Storage) read(ctx context.Context, path string, w io.Writer, opt pairStorageRead) (n int64, err error) {
	panic("not implemented")
}
