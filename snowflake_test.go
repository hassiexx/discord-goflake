package dgoflake

import (
	"strconv"
	"testing"
	"time"
)

func TestSnowflake_Increment(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := id & 0xFFF
	actual := snowflake.Increment()
	if expected != actual {
		t.Errorf("Snowflake increment was not correct for %d, expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_Int64(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := int64(id)
	actual := snowflake.Int64()
	if expected != actual {
		t.Errorf("Snowflake int64 was not correct for %d, expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_InternalProcessID(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := (id & 0x1F000) >> 12
	actual := snowflake.InternalProcessID()
	if expected != actual {
		t.Errorf("Snowflake internal process ID was not correct for %d, expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_InternalWorkerID(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := (id & 0x3E0000) >> 17
	actual := snowflake.InternalWorkerID()
	if expected != actual {
		t.Errorf("Snowflake internal worker ID was not correct for %d, expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_String(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := strconv.FormatUint(id, 10)
	actual := snowflake.String()
	if expected != actual {
		t.Errorf("Snowflake string was not correct for %d, expected %s, got %s", id, expected, actual)
	}
}

func TestSnowflake_Timestamp(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := time.Unix(0, ((int64(id)>>22)+int64(DiscordEpoch))*1000000).String()
	actual := snowflake.Timestamp().String()
	if expected != actual {
		t.Errorf("Snowflake timestamp was not correct for %d, expected %s, got %s", id, expected, actual)
	}
}

func TestSnowflake_UInt64(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := id
	actual := snowflake.UInt64()
	if expected != actual {
		t.Errorf("Snowflake uint64 was not correct for %d, expected %d, got %d", id, expected, actual)
	}
}

func BenchmarkSnowflake_Timestamp(b *testing.B) {
	snowflake := ParseInt(577654653257383975)
	for i := 0; i < b.N; i++ {
		snowflake.Timestamp()
	}
}

