package tests

import (
	"os"
	"testing"

	"github.com/ragulmathawa/go-storage/tests"
)

func TestStorage(t *testing.T) {
	if os.Getenv("STORAGE_COS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_COS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestStorager(t, setupTest(t))
}

func TestMultiparter(t *testing.T) {
	if os.Getenv("STORAGE_COS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_COS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestMultiparter(t, setupTest(t))
}

func TestDir(t *testing.T) {
	if os.Getenv("STORAGE_COS_INTEGRATION_TEST") != "on" {
		t.Skipf("STORAGE_COS_INTEGRATION_TEST is not 'on', skipped")
	}
	tests.TestDirer(t, setupTest(t))
}
