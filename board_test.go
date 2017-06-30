package main

import (
	"testing"
	"fmt"
)

func TestCopy(t *testing.T) {
	b := board{boardLine{e, x, e}, boardLine{e, e, e}, boardLine{e, o, e}}
	if b.diff(b.copy()) != 0 {
		fmt.Println(b, b.copy())
		t.Fail()
	}
}

func TestRotate(t *testing.T) {
	b    := board{boardLine{e, e, e}, boardLine{e, e, e}, boardLine{e, e, e}}
	bRot := board{boardLine{e, e, e}, boardLine{e, e, e}, boardLine{e, e, e}}
	if b.rotate().diff(bRot) != 0 {
		t.Fail()
	}

	b      = board{boardLine{e, x, e}, boardLine{e, e, e}, boardLine{e, o, e}}
	bRot1 := board{boardLine{e, e, e}, boardLine{o, e, x}, boardLine{e, e, e}}
	bRot2 := board{boardLine{e, o, e}, boardLine{e, e, e}, boardLine{e, x, e}}
	bRot3 := board{boardLine{e, e, e}, boardLine{x, e, o}, boardLine{e, e, e}}
	if b.rotate().diff(bRot1) != 0 {
		t.Fail()
	}
	if b.rotate().rotate().diff(bRot2) != 0 {
		t.Fail()
	}
	if b.rotate().rotate().rotate().diff(bRot3) != 0 {
		t.Fail()
	}
	if b.rotate().rotate().rotate().rotate().diff(b) != 0 {
		t.Fail()
	}
}
