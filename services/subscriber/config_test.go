package subscriber_test

import (
	"testing"

	"github.com/influxdata/config"
	"github.com/influxdb/influxdb/services/subscriber"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c subscriber.Config
	if err := config.Decode(`
enabled = false
`, &c); err != nil {
		t.Fatal(err)
	}

	// Validate configuration.
	if c.Enabled != false {
		t.Fatalf("unexpected enabled state: %v", c.Enabled)
	}
}
