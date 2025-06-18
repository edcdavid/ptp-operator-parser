package parsers

import (
	"fmt"
	"regexp"
	"strconv"
)

type Ptp4lLog struct {
	Timestamp    float64
	Config       string
	Severity     *int
	RMS          *int
	Max          *int
	Freq         int
	FreqError    *int
	Delay        int
	DelayError   *int
	MasterOffset *int
	Servo        *string
}

var ptp4lRegex = regexp.MustCompile(
	`^ptp4l\[(?P<timestamp>\d+\.\d+)\]: \[(?P<config>ptp4l\.\d+\.config)(?::(?P<severity>\d+))?\] ` +
		`(?:(?:rms (?P<rms>\d+) max (?P<max>\d+) )?freq (?P<freq>[-+]?\d+)(?: \+/- (?P<freqError>\d+))? delay (?P<delay>\d+)(?: \+/- (?P<delayError>\d+))?|` +
		`master offset (?P<masterOffset>[-+]?\d+) (?P<servo>\w+) freq (?P<freq>[-+]?\d+) path delay (?P<delay>\d+))`)

func parsePtp4lLine(line string) (*Ptp4lLog, error) {
	match := ptp4lRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, fmt.Errorf("failed to parse line: %q", line)
	}

	result := &Ptp4lLog{}
	groupNames := ptp4lRegex.SubexpNames()

	for i, name := range groupNames {
		if name == "" {
			continue
		}
		val := match[i]
		if val == "" {
			continue
		}

		switch name {
		case "timestamp":
			parsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid timestamp %q in line %q: %w", val, line, err)
			}
			result.Timestamp = parsed

		case "config":
			result.Config = val

		case "severity":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid severity %q in line %q: %w", val, line, err)
			}
			result.Severity = &parsed

		case "rms":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid rms %q in line %q: %w", val, line, err)
			}
			result.RMS = &parsed

		case "max":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid max %q in line %q: %w", val, line, err)
			}
			result.Max = &parsed

		case "freq":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid freq %q in line %q: %w", val, line, err)
			}
			result.Freq = parsed

		case "freqError":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid freq error %q in line %q: %w", val, line, err)
			}
			result.FreqError = &parsed

		case "delay":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid delay %q in line %q: %w", val, line, err)
			}
			result.Delay = parsed

		case "delayError":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid delay error %q in line %q: %w", val, line, err)
			}
			result.DelayError = &parsed

		case "masterOffset":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid master offset %q in line %q: %w", val, line, err)
			}
			result.MasterOffset = &parsed

		case "servo":
			result.Servo = &val
		}
	}

	return result, nil
}
