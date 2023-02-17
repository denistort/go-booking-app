package app

import (
	"github.com/denistort/go-booking-app/cmd/web/mock"
	"testing"
)

func TestApp(t *testing.T) {
	app, err := New(mock.Config)

	// First Case should not be error if you try to create app
	if err != nil {
		t.Error("failed on creation")
	}
	// Second Case should be error if you try to created app twice
	if _, e := New(mock.Config); e == nil {
		t.Error("should be error if you try to created app twice")
	}
	// Third Case should not be error when you try to start app
	if startedError := app.StartServer(); startedError == nil {
		t.Error("should not be error when you try to start app")
	}
}
