package loop_test

import "testing"

func TestLoop(t *testing.T) {
	n := 0
	for n < 5 { //等同于传统  while(n<5)
		t.Log(n)
		n++
	}
}
