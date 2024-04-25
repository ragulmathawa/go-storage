package tests

import (
	"testing"

	"github.com/ragulmathawa/go-storage/types"

	"github.com/beyondstorage/go-storage/services/memory"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for memory")

	store, err := memory.NewStorager()
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
