package test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) //false
	t.Log(a == d) //true
}
