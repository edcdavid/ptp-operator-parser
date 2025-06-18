package parsers

import (
	"reflect"
	"testing"
)

func TestParseTs2PhcLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *Ts2PhcLog
		wantErr  bool
	}{
		{
			name: "ens2f1 with master",
			line: "ts2phc[82674.465]: [ts2phc.0.config] ens2f1 master offset 0 s2 freq -0",
			expected: &Ts2PhcLog{
				Process:   "ts2phc",
				Uptime:    82674.465,
				Config:    "ts2phc.0.config",
				Interface: "ens2f1",
				Offset:    0,
				Servo:     "s2",
				Freq:      0,
				Extra:     "",
			},
		},
		{
			name: "ptp0 with master",
			line: "ts2phc[82674.465]: [ts2phc.0.config] /dev/ptp0 master offset 0 s2 freq -0",
			expected: &Ts2PhcLog{
				Process:   "ts2phc",
				Uptime:    82674.465,
				Config:    "ts2phc.0.config",
				Interface: "/dev/ptp0",
				Offset:    0,
				Servo:     "s2",
				Freq:      0,
				Extra:     "",
			},
		},
		{
			name: "ptp4 with master",
			line: "ts2phc[82674.465]: [ts2phc.0.config] /dev/ptp4 master offset 0 s2 freq -0",
			expected: &Ts2PhcLog{
				Process:   "ts2phc",
				Uptime:    82674.465,
				Config:    "ts2phc.0.config",
				Interface: "/dev/ptp4",
				Offset:    0,
				Servo:     "s2",
				Freq:      0,
				Extra:     "",
			},
		},
		{
			name: "ptp4 without master",
			line: "ts2phc[82674.465]: [ts2phc.0.config] /dev/ptp4 offset 0 s2 freq -0",
			expected: &Ts2PhcLog{
				Process:   "ts2phc",
				Uptime:    82674.465,
				Config:    "ts2phc.0.config",
				Interface: "/dev/ptp4",
				Offset:    0,
				Servo:     "s2",
				Freq:      0,
				Extra:     "",
			},
		},
		{
			name: "ptp6 with extra 'holdover'",
			line: "ts2phc[82674.465]: [ts2phc.0.config] /dev/ptp6 offset 0 s3 freq +0 holdover",
			expected: &Ts2PhcLog{
				Process:   "ts2phc",
				Uptime:    82674.465,
				Config:    "ts2phc.0.config",
				Interface: "/dev/ptp6",
				Offset:    0,
				Servo:     "s3",
				Freq:      0,
				Extra:     "holdover",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTs2PhcLine(tt.line)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("parsed result mismatch:\ngot:  %+v\nwant: %+v", got, tt.expected)
			}
		})
	}
}
