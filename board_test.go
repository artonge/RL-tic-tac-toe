package main

import "testing"

func TestIsEqual(t *testing.T) {
	b1 := board{boardLine{e, e, e}, boardLine{e, e, e}, boardLine{e, e, e}}
	b2 := board{boardLine{e, e, e}, boardLine{e, e, e}, boardLine{e, e, e}}
	if !b1.isEqual(b2) {
		t.Fail()
	}

	b1 = board{boardLine{e, x, e}, boardLine{e, e, e}, boardLine{e, o, e}}
	b2 = board{boardLine{e, x, e}, boardLine{e, e, e}, boardLine{e, o, e}}
	if !b1.isEqual(b2) {
		t.Fail()
	}
}
