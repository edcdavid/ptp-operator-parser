
package parsers

import (
	"fmt"
	"regexp"
	"strconv"
)

type GnssLog struct {
	Process    string
	Uptime     int64
	Config     string
	Interface  string
	Status     int
	Offset     int
	Servo      string
}

var gnssRegex = regexp.MustCompile(`(?P<process>\w+)\[(?P<uptime>\d+)\]:\[(?P<config>[^\]]+)\] (?P<iface>\S+) gnss_status (?P<status>\d+) offset (?P<offset>-?\d+) (?P<servo>s\d+)`)

func parseGnssLine(line string) (*GnssLog, error) {
	match := gnssRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, fmt.Errorf("failed to parse line: %q", line)
	}

	result := &GnssLog{}
	groupNames := gnssRegex.SubexpNames()

	for i, name := range groupNames {
		if name == "" {
			continue
		}
		val := match[i]
		switch name {
		case "process":
			result.Process = val
		case "uptime":
			parsed, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid uptime %q: %w", val, err)
			}
			result.Uptime = parsed
		case "config":
			result.Config = val
		case "iface":
			result.Interface = val
		case "status":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid status %q: %w", val, err)
			}
			result.Status = parsed
		case "offset":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid offset %q: %w", val, err)
			}
			result.Offset = parsed
		case "servo":
			result.Servo = val
		}
	}

	return result, nil
}
