package mask

import (
	"fmt"
)

type (
	// Rule represents a masking rule.
	Rule struct {
		maskType maskType
	}
	// Rules is a collection of Rule.
	Rules []Rule

	// maskType is a method of masking.
	maskType int
)

const (
	maskTypeUnknown = iota
	maskTypeRaw
	maskTypeHash
	maskTypeChecksum
)

// NewRules returns a new rules from a slice of rules.
func NewRules(ss []string) (Rules, error) {
	rules := make(Rules, len(ss))
	for i, s := range ss {
		t, err := toMaskType(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert mask type at %d: %v", i, err)
		}
		rules[i] = Rule{
			maskType: t,
		}
	}
	return rules, nil
}

// Convert converts s according to the rule.
func (r *Rule) Convert(s string) (string, error) {
	switch r.maskType {
	case maskTypeRaw:
		return s, nil
	case maskTypeHash:
		return hash(s), nil
	case maskTypeChecksum:
		return checksum(s), nil
	default:
		return s, fmt.Errorf("unknown mask type: %v", r.maskType)
	}
}

func toMaskType(s string) (maskType, error) {
	switch s {
	case "Raw":
		return maskTypeRaw, nil
	case "Hash":
		return maskTypeHash, nil
	case "Checksum":
		return maskTypeChecksum, nil
	default:
		return maskTypeUnknown, fmt.Errorf("unknown mask type: %s", s)
	}
}
