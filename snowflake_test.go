package dgoflake

import (
	"testing"
	"time"
)

func TestSnowflake_Increment(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := id & 0xFFF
	actual := snowflake.Increment()
	if expected != actual {
		t.Errorf("unexpected increment value for snowflake %d: expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_Int64(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := int64(id)
	actual := snowflake.Int64()
	if expected != actual {
		t.Errorf("unexpected int64 value for snowflake %d: expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_InternalProcessID(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := (id & 0x1F000) >> 12
	actual := snowflake.InternalProcessID()
	if expected != actual {
		t.Errorf("unexpected internal process ID value for snowflake %d: expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_InternalWorkerID(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := (id & 0x3E0000) >> 17
	actual := snowflake.InternalWorkerID()
	if expected != actual {
		t.Errorf("unexpected internal worker ID value for snowflake %d: expected %d, got %d", id, expected, actual)
	}
}

func TestSnowflake_String(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := "577645285396840449"
	actual := snowflake.String()
	if expected != actual {
		t.Errorf("unexpected string value for snowflake %d: expected %s, got %s", id, expected, actual)
	}
}

func TestSnowflake_Timestamp(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := time.Unix(0, ((int64(id)>>22)+int64(DiscordEpoch))*1000000).String()
	actual := snowflake.Timestamp().String()
	if expected != actual {
		t.Errorf("unexpected timestamp value for snowflake %d: expected %s, got %s", id, expected, actual)
	}
}

func TestSnowflake_UInt64(t *testing.T) {
	var id uint64 = 577645285396840449
	snowflake := ParseInt(id)
	expected := id
	actual := snowflake.UInt64()
	if expected != actual {
		t.Errorf("unexpected uint64 value for snowflake %d: expected %d, got %d", id, expected, actual)
	}
}

func BenchmarkSnowflake_Timestamp(b *testing.B) {
	snowflake := ParseInt(577654653257383975)
	for i := 0; i < b.N; i++ {
		snowflake.Timestamp()
	}
}
