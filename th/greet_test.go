package th

import "testing"

func TestGreet(t *testing.T) {
	want := "สวัสดี-Hi There"

	if got := Greet(); got != want {
		t.Errorf("Hello() = %q, want: %q", got, want)
	}
}
