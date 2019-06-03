package th

import "gitlab.com/yuttasakyo/quote"

// Greet : "สวัสดี-Hi There"
func Greet() string {
	return "สวัสดี-" + quote.Speak()
}
