package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q bug got %q", expected, repeated)
	}
}

func TestRepeatTimes(t *testing.T) {
	repeated := RepeatTimes("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q bug got %q", expected, repeated)
	}
}
