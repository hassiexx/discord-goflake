package dgoflake

import (
	"strconv"

	"golang.org/x/xerrors"
)

// ParseInt parses an int to a snowflake.
func ParseInt(id uint64) *Snowflake {
	return &Snowflake{id}
}

// ParseString parses a string to a snowflake.
func ParseString(id string) (*Snowflake, error) {
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, xerrors.Errorf("dgoflake: failed to parse string to snowflake: %w", err)
	}
	return &Snowflake{i}, nil
}

// ParseEpochMilli parses an epoch milliseconds timestamp to a snowflake.
func ParseEpochMilli(epoch uint64) (*Snowflake, error) {
	if epoch < DiscordEpoch {
		return nil, xerrors.New("dgoflake: epoch cannot be earlier than the Discord epoch")
	}
	return &Snowflake{(epoch - DiscordEpoch) << 22}, nil
}

// ParseEpochSec parses an epoch seconds timestamp to a snowflake.
func ParseEpochSec(epoch uint64) (*Snowflake, error) {
	return ParseEpochMilli(epoch * 1000)
}
