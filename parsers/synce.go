
package parsers

import (
	"fmt"
	"regexp"
	"strconv"
)

type SynceLog struct {
	Process        string
	Uptime         float64
	Config         string
	Interface      string
	Device         string
	EecState       string
	ClockQuality   string
	ExtQl          *int 
	Ql             *int
	NetworkOption  int
	Servo          string
}
var synceRegex = regexp.MustCompile(`(?P<process>\w+)\[(?P<uptime>[\d.]+)\]: \[(?P<config>[^\]]+)\] (?P<iface>\S+)(?: clock_quality (?P<clock_quality>\S+))? device (?P<device>\S+)(?: eec_state (?P<eec_state>\S+))?(?: ext_ql (?P<ext_ql>0x[\da-fA-F]+))?(?: network_option (?P<network_option>\d+))(?: ql (?P<ql>0x[\da-fA-F]+))? (?P<servo>s\d+)`)

func parseSynceLine(line string) (*SynceLog, error) {
	match := synceRegex.FindStringSubmatch(line)
	if match == nil {
		return nil, fmt.Errorf("failed to parse line: %q", line)
	}

	result := &SynceLog{}
	groupNames := synceRegex.SubexpNames()

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
		case "clock_quality":
			result.ClockQuality = val
		case "device":
			result.Device = val
		case "eec_state":
			result.EecState = val
		case "ext_ql":
			if val != "" {
				n, err := strconv.ParseInt(val, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid ext_ql %q: %w", val, err)
				}
				tmp := int(n)
				result.ExtQl = &tmp
			}
		case "ql":
			if val != "" {
				n, err := strconv.ParseInt(val, 0, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid ql %q: %w", val, err)
				}
				tmp := int(n)
				result.Ql = &tmp
			}
		case "network_option":
			parsed, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("invalid network_option: %w", err)
			}
			result.NetworkOption = parsed
		case "servo":
			result.Servo = val
		}
	}

	return result, nil
}
