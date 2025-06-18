package parsers

import (
	"reflect"
	"testing"
)

func strPtr(s string) *string { return &s }

func TestParsePtp4lLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *Ptp4lLog
		wantErr  bool
	}{
		{
			name: "rms/max/freq/delay with errors",
			line: "ptp4l[74737.942]: [ptp4l.0.config] rms 53 max 74 freq -16642 +/- 40 delay 1089 +/- 20",
			expected: &Ptp4lLog{
				Timestamp:  74737.942,
				Config:     "ptp4l.0.config",
				RMS:        intPtr(53),
				Max:        intPtr(74),
				Freq:       -16642,
				FreqError:  intPtr(40),
				Delay:      1089,
				DelayError: intPtr(20),
			},
		},
		{
			name: "master offset with servo and path delay",
			line: "ptp4l[365195.391]: [ptp4l.0.config] master offset -1 s2 freq -3972 path delay 89",
			expected: &Ptp4lLog{
				Timestamp:    365195.391,
				Config:       "ptp4l.0.config",
				MasterOffset: intPtr(-1),
				Servo:        strPtr("s2"),
				Freq:         -3972,
				Delay:        89,
			},
		},
		{
			name: "severity field included",
			line: "ptp4l[5196755.139]: [ptp4l.0.config:6] rms 53 max 74 freq -16642 +/- 40 delay 1089 +/- 20",
			expected: &Ptp4lLog{
				Timestamp:  5196755.139,
				Config:     "ptp4l.0.config",
				Severity:   intPtr(6),
				RMS:        intPtr(53),
				Max:        intPtr(74),
				Freq:       -16642,
				FreqError:  intPtr(40),
				Delay:      1089,
				DelayError: intPtr(20),
			},
		},
		{
			name:    "invalid line should error",
			line:    "this is not a ptp4l log",
			wantErr: true,
		},
		{
			name:    "malformed freq should error",
			line:    "ptp4l[74737.942]: [ptp4l.0.config] rms 53 max 74 freq notanumber +/- 40 delay 1089 +/- 20",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parsePtp4lLine(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePtp4lLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("parsePtp4lLine() = %+v,\nexpected %+v", got, tt.expected)
			}
		})
	}
}
