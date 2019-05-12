package dgoflake

import (
	"strconv"
	"time"
)

// DiscordEpoch is Discord's epoch in milliseconds.
const DiscordEpoch uint64 = 1420070400000

// Snowflake is the type for a Discord snowflake.
type Snowflake struct {
	id uint64
}

// Int64 gets the snowflake as an int64.
func (s *Snowflake) Int64() int64 {
	return int64(s.id)
}

// Increment gets the generated number on the process for the ID.
func (s *Snowflake) Increment() uint64 {
	return s.id & 0xFFF
}

// InternalProcessID gets the internal process ID for this snowflake.
func (s *Snowflake) InternalProcessID() uint64 {
	return (s.id & 0x1F000) >> 12
}

// InternalWorkerID gets the internal worker ID for this snowflake.
func (s *Snowflake) InternalWorkerID() uint64 {
	return (s.id & 0x3E0000) >> 17
}

// IsValid checks whether the snowflake is valid.
func (s *Snowflake) IsValid() bool {
	return (s.id>>22)+DiscordEpoch >= DiscordEpoch
}

// String gets the snowflake as a string.
func (s *Snowflake) String() string {
	return strconv.FormatInt(s.Int64(), 10)
}

// UInt64 gets the snowflake as a uint64.
func (s *Snowflake) UInt64() uint64 {
	return s.id
}

// Timestamp converts the snowflake to a timestamp.
func (s *Snowflake) Timestamp() time.Time {
	return time.Unix(0, ((int64(s.id)>>22)+int64(DiscordEpoch))*1000000)
}
