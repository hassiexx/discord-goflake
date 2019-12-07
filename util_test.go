package dgoflake

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	var id uint64 = 577637225211101206
	snowflake := ParseInt(id)
	if snowflake.UInt64() != id {
		t.Errorf("failed to parse snowflake from int %d", id)
	}
}

func TestParseString(t *testing.T) {
	id := "577452678049955841"
	snowflake, err := ParseString(id)
	if err != nil {
		t.Errorf("failed to parse snowflake from string %s: %v", id, err)
	} else if snowflake.String() != id {
		t.Errorf("failed to parse snowflake from string %s: got %s", id, snowflake.String())
	}
}

func TestParseStringError(t *testing.T) {
	id := "57745267804995584X"
	_, err := ParseString(id)
	if err == nil {
		t.Errorf("expected error for invalid id")
	}
}

func TestParseEpochMilli(t *testing.T) {
	var id uint64 = 576769138849218561
	var epoch uint64 = 1557582878554
	expected := ParseInt(id)
	actual, err := ParseEpochMilli(epoch)
	if err != nil {
		t.Errorf("failed to parse snowflake from epoch ms %d: %v", epoch, err)
	} else if expected.Timestamp().String() != actual.Timestamp().String() {
		t.Errorf("failed to parse snowflake from epoch ms %d: expected %s, got %s", epoch, expected.Timestamp().String(),
			actual.Timestamp().String())
	}
}

func TestParseEpochMilliEarlier(t *testing.T) {
	var epoch uint64 = 1420070300000
	_, err := ParseEpochMilli(epoch)
	if err == nil {
		t.Errorf("expected error for epoch being earlier than Discord's epoch")
	}
}

func TestParseEpochSec(t *testing.T) {
	var epoch uint64 = 1557793289
	snowflake, err := ParseEpochSec(epoch)
	if err != nil {
		t.Errorf("failed to parse snowflake from epoch sec %d: %v", epoch, err)
	} else if int64(epoch) != snowflake.Timestamp().Unix() {
		t.Errorf("failed to parse snowflake from epoch s %d: expected %d, got %d", epoch, epoch,
			snowflake.Timestamp().Unix())
	}
}

func BenchmarkParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseInt(577654653257383975)
	}
}

func BenchmarkParseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseString("577654653257383975")
	}
}

func BenchmarkParseEpochMilli(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseEpochMilli(1557582878554)
	}
}

func BenchmarkParseEpochSec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseEpochSec(1557793289)
	}
}
