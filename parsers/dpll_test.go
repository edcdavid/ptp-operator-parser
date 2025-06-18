package parsers

import (
	"reflect"
	"testing"
)

func TestParseDpllLine(t *testing.T) {
	line := "dpll[1700598434]:[ts2phc.0.config] ens2f0 frequency_status 3 offset 0 phase_status 3 pps_status 1 s2"

	expected := &DpllLog{
		Process:         "dpll",
		Uptime:          1700598434,
		Config:          "ts2phc.0.config",
		Interface:       "ens2f0",
		FrequencyStatus: 3,
		Offset:          0,
		PhaseStatus:     3,
		PpsStatus:       1,
		Servo:           "s2",
	}

	got, err := parseDpllLine(line)
	if err != nil {
		t.Fatalf("parseDpllLine() error: %v", err)
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("parseDpllLine() = %+v, want %+v", got, expected)
	}
}
