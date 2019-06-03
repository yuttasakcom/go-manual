package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello"

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want: %q", got, want)
	}
}

func TestThHello(t *testing.T) {
	want := "สวัสดี-Hi There"

	if got := ThHello(); got != want {
		t.Errorf("ThHello() = %q, want: %q", got, want)
	}
}
