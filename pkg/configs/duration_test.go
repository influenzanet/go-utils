package configs

import (
	"encoding/json"
	"testing"
	"time"
)

type testDuration struct {
	Value Duration `json:"value"`
}

func testUnmarshall(t *testing.T, jsonString string, expect time.Duration) {
	var d testDuration
	err := json.Unmarshal([]byte(jsonString), &d)
	if err != nil {
		t.Fatalf("Unable to parse duration '%s' : %s", jsonString, err)
	}
	if d.Value.Duration != expect {
		t.Fatalf("Invalid duration expecting '%s', got '%s'", expect, d.Value.Duration)
	}
}

func TestUnmarshallDuration(t *testing.T) {
	testUnmarshall(t, `{"value":"2h"}`, 2*time.Hour)
	testUnmarshall(t, `{"value":"3s"}`, 3*time.Second)
}

func testParseDuration(t *testing.T, value string, defaultUnit string, expect time.Duration) {
	d, err := ParseDuration(value, defaultUnit)
	if err != nil {
		t.Fatalf("Unable to parse duration '%s' : %s", value, err)
	}
	if d != expect {
		t.Fatalf("Invalid duration expecting '%s', got '%s'", expect, d)
	}
}

func TestParseDuration(t *testing.T) {
	testParseDuration(t, "2", "s", 2*time.Second)
	testParseDuration(t, "3m", "s", 3*time.Minute)
	testParseDuration(t, "4", "h", 4*time.Hour)
	testParseDuration(t, "5h", "h", 5*time.Hour)
}
