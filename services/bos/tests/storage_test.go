package tests

import (
	"os"
	"testing"

	"github.com/ragulmathawa/go-storage/tests"
)

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_BOS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_BOS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}
