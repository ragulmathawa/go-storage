package tests

import (
	"os"
	"testing"

	azblob "github.com/beyondstorage/go-storage/services/azblob/v3"
	"github.com/google/uuid"
	ps "github.com/ragulmathawa/go-storage/pairs"
	"github.com/ragulmathawa/go-storage/types"
)

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for azblob")

	store, err := azblob.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_AZBLOB_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_AZBLOB_NAME")),
		ps.WithEndpoint(os.Getenv("STORAGE_AZBLOB_ENDPOINT")),
		ps.WithWorkDir("/"+uuid.New().String()+"/"),
		ps.WithEnableVirtualDir(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
