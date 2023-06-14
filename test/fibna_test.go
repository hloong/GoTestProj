package const_test

import "testing"

func TestFibna(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	t.Log(" ", a)
	for i := 0; i < 8; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchangeNomarl(t *testing.T) {
	a := 1
	b := 2
	temp := a
	a = b
	b = temp
	t.Log(a, b)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
