package clock_test

import (
	"github.com/bencicandrej/tricks/clock"
	"testing"
	"time"
)

func TestBrokenClock(t *testing.T) {
	expectedTime, err := time.Parse("2006-01-02", "2016-12-15")
	if err != nil {
		t.Fatalf("parse time: %v", err)
	}

	clockInstance := clock.NewBrokenClock(expectedTime)

	if now := clockInstance.Now(); now != expectedTime {
		t.Errorf("clock.Now() = %q, wanted %q", now, expectedTime)
	}
}
