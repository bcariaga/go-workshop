package main

import "testing"

func TestProccessIncoming(t *testing.T) {
	var in float32 = 1000
	var want float32 = 2000

	got := ProccessIncome(in)

	if got != want {
		t.Errorf("ProccessIncome(%v) == %v, want %v", in, got, want)
	}
}
