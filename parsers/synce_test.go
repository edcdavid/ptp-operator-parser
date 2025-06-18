package parsers

import (
	"reflect"
	"testing"
)


func intPtr(i int) *int { return &i }

func TestParseSynceLine(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *SynceLog
		wantErr  bool
	}{
		{
			name: "eec_state line",
			line: "synce4l[5196755.139]: [synce4l.0.config] ens7f0 device synce1 eec_state EEC_HOLDOVER network_option 2 s1",
			expected: &SynceLog{
				Process:       "synce4l",
				Uptime:        5196755.139,
				Config:        "synce4l.0.config",
				Interface:     "ens7f0",
				Device:        "synce1",
				EecState:      "EEC_HOLDOVER",
				NetworkOption: 2,
				Servo:         "s1",
			},
		},
		{
			name: "clock_quality line with hex ql",
			line: "synce4l[5196755.139]: [synce4l.0.config] ens7f0 clock_quality PRTC device synce1 ext_ql 0x20 network_option 2 ql 0x1 s2",
			expected: &SynceLog{
				Process:       "synce4l",
				Uptime:        5196755.139,
				Config:        "synce4l.0.config",
				Interface:     "ens7f0",
				ClockQuality:  "PRTC",
				Device:        "synce1",
				ExtQl:         intPtr(0x20),
				NetworkOption: 2,
				Ql:            intPtr(0x1),
				Servo:         "s2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSynceLine(tt.line)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("parseSynceLine() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}
