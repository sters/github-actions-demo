package main

import "testing"

func Test_say(t *testing.T) {
	want := "I am bot, hello!"
	got := say("bot", "hello!")

	if want != got {
		t.Errorf("want = %s, got = %s", want, got)
	}
}
