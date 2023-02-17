package appRouter

import (
	"github.com/denistort/go-booking-app/cmd/web/mock"
	"testing"
)

func TestAppRouter(t *testing.T) {
	r := New(mock.Config)

	// First Case should not be error if you try to create app
	if r == nil {
		t.Error("failed on creation app router")
	}

	if handler := r.Start(); handler == nil {
		t.Error("failed on start app router")
	}
}
