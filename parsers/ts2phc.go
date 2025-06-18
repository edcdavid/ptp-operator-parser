package parsers

import (
	"fmt"
	"regexp"
	"strconv"
)

type Ts2PhcLog struct {
	Process   string
	Uptime    float64
	Config    string
	Interface string
	Offset    int
	Servo     string
	Freq      int
	Extra     string
}

var ts2phcRegex = regexp.MustCompile(`(?P<process>\w+)\[(?P<uptime>[\d.]+)\]: \[(?P<config>[^\]]+)\] (?P<iface>\S+) (?:master )?offset (?P<offset>-?\d+) (?P<servo>s\d+) freq (?P<freq>[+-]?\d+)(?: (?P<extra>\w+))?`)

func parseTs2PhcLine(line string) (*Ts2PhcLog, error) {
	match := ts2phcRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, fmt.Errorf("failed to parse line: %q", line)
	}

	result := &Ts2PhcLog{}
	groupNames := ts2phcRegex.SubexpNames()

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
				return nil, fmt.Errorf("invalid uptime %q in line %q: %w", val, line, err)
			}
			result.Uptime = parsed
		case "config":
			result.Config = val
		case "iface":
			result.Interface = val
		case "offset":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid offset %q in line %q: %w", val, line, err)
			}
			result.Offset = parsed
		case "servo":
			result.Servo = val
		case "freq":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid frequency offset %q in line %q: %w", val, line, err)
			}
			result.Freq = parsed
		case "extra":
			result.Extra = val
		}
	}

	return result, nil
}

