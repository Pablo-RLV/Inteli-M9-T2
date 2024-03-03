package tests

import (
	common "Ponderada4/common"
	"testing"
	"os"
)

func TestEnv(t *testing.T) {
	common.LoadEnv()
	if os.Getenv("BROKER_ADDR") == "" {
		t.Errorf("BROKER_ADDR not set")
	}
	if os.Getenv("HIVE_USER") == "" {
		t.Errorf("HIVE_USER not set")
	}
	if os.Getenv("HIVE_PSWD") == "" {
		t.Errorf("HIVE_PSWD not set")
	}
}