package config

import (
	"github.com/denistort/go-booking-app/cmd/web/mock"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New(mock.Config)

	// First Case should not be error if you try to create config
	if err != nil {
		t.Error("failed on creation config")
	}
	if config == nil {
		t.Error("should be config")
	}
	// Second Case should be error if you try to created config twice
	if _, e := New(mock.Config); e == nil {
		t.Error("should be error if you try to created app twice")
	}
}
