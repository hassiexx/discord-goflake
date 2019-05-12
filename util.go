package dgoflake

import (
	"strconv"
)

// ParseInt parses an int to a snowflake.
func ParseInt(id int) (*Snowflake, error) {
	if id < 0 {
		return nil, newError("ID must be greater than 0", nil)
	}

	return &Snowflake{uint64(id)}, nil
}

// ParseString parses a string to a snowflake.
func ParseString(id string) (*Snowflake, error) {
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, newError("Failed to parse string:", err)
	}
	return &Snowflake{i}, nil
}

// ParseEpochMilli parses a epoch milliseconds timestamp to a snowflake.
func ParseEpochMilli(epoch uint64) (*Snowflake, error) {
	if epoch < DiscordEpoch {
		return nil, newError("Epoch milliseconds timestamp cannot be less than the Discord epoch", nil)
	}

	return &Snowflake{(epoch - DiscordEpoch) << 22}, nil
}
