package tests

import (
	common "Ponderada7/common"
	"os"
	"testing"
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
