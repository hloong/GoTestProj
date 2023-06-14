package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open    = 3 << iota //向左移位，相当于 3*2的次方
	Close               //3*2^1  累加
	Pending             //3*2^2
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	t.Log(Open, Close, Pending)
}
