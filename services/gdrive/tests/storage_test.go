package tests

import (
	"os"
	"testing"

	"github.com/ragulmathawa/go-storage/tests"
)

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_GDRIVE_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_GDRIVE_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}

func TestCopier(t *testing.T) {
	if os.Getenv("STORAGE_GDRIVE_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_GDRIVE_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestCopier(t, setupTest(t))
}
