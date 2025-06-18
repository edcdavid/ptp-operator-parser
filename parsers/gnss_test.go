package parsers

import (
	"reflect"
	"testing"
)

func TestParseGnssLine(t *testing.T) {
	line := "gnss[1689014431]:[ts2phc.0.config] ens2f1 gnss_status 5 offset 0 s0"
	expected := &GnssLog{
		Process:   "gnss",
		Uptime:    1689014431,
		Config:    "ts2phc.0.config",
		Interface: "ens2f1",
		Status:    5,
		Offset:    0,
		Servo:     "s0",
	}

	got, err := parseGnssLine(line)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("parseGnssLine() = %+v, want %+v", got, expected)
	}
}
