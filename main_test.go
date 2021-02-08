package main

import "testing"

func Test_say(t *testing.T) {
	want := "I'm bot, hello!"
	got := say("hello!")

	if want != got {
		t.Errorf("want = %s, got = %s", want, got)
	}
}
