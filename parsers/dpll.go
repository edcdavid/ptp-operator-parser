
package parsers

import (
	"fmt"
	"regexp"
	"strconv"
)

type DpllLog struct {
	Process         string
	Uptime          float64
	Config          string
	Interface       string
	FrequencyStatus int
	Offset          int
	PhaseStatus     int
	PpsStatus       int
	Servo           string
}
var dpllRegex = regexp.MustCompile(`(?P<process>\w+)\[(?P<uptime>\d+(?:\.\d+)?)\]:\[(?P<config>[^\]]+)\] (?P<iface>\S+) frequency_status (?P<freq_status>\d+) offset (?P<offset>-?\d+) phase_status (?P<phase_status>\d+) pps_status (?P<pps_status>\d+) (?P<servo>s\d+)`)

func parseDpllLine(line string) (*DpllLog, error) {
	match := dpllRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, fmt.Errorf("failed to parse line: %q", line)
	}

	result := &DpllLog{}
	groupNames := dpllRegex.SubexpNames()

	for i, name := range groupNames {
		if name == "" {
			continue
		}
		val := match[i]
		switch name {
		case "process":
			result.Process = val
		case "uptime":
			parsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid uptime: %w", err)
			}
			result.Uptime = parsed
		case "config":
			result.Config = val
		case "iface":
			result.Interface = val
		case "freq_status":
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid frequency_status: %w", err)
			}
			result.FrequencyStatus = n
		case "offset":
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid offset: %w", err)
			}
			result.Offset = n
		case "phase_status":
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid phase_status: %w", err)
			}
			result.PhaseStatus = n
		case "pps_status":
			n, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid pps_status: %w", err)
			}
			result.PpsStatus = n
		case "servo":
			result.Servo = val
		}
	}

	return result, nil
}
