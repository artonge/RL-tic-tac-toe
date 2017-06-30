package main

import (
	"testing"
	"fmt"
)


func TestIncrementalMean(t *testing.T) {
	mean := 0.0
	k := 0

	mean, k = incrementalMean(mean, 10, k)
	if mean != 10 || k != 1 {
		fmt.Println(mean, k)
		t.Fail()
	}

	mean, k = incrementalMean(mean, 5, k)
	if mean != 7.5 || k != 2 {
		fmt.Println(mean, k)
		t.Fail()
	}

	mean, k = incrementalMean(mean, 15, k)
	if mean != 10 || k != 3 {
		fmt.Println(mean, k)
		t.Fail()
	}
}
