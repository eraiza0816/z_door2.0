package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := loadConfig()
	if config.UnlockAngle == 0 {
		t.Errorf("Expected UnlockAngle to be non-zero")
	}
	if config.LockAngle == 0 {
		t.Errorf("Expected LockAngle to be non-zero")
	}
	if config.SleepSec == 0 {
		t.Errorf("Expected SleepSec to be non-zero")
	}
	if config.AutoLockSec == 0 {
		t.Errorf("Expected AutoLockSec to be non-zero")
	}
}

func TestLoadUsers(t *testing.T) {
	users := loadUsers()
	if len(users) == 0 {
		t.Errorf("Expected users to be non-empty")
	}
}

func TestGetIdm(t *testing.T) {
	text := "ID=1234 "
	idm := getIdm(text)
	if idm != "1234" {
		t.Errorf("Expected idm to be 1234, got %s", idm)
	}
}
