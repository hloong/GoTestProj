package const_test

import (
	"testing"
)

type MyInt int64

func TestType(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c) //go语言的类型转换比较严格
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1//go指针不支持运算,这样会报错
	t.Log(a, aPtr)           //1 0xc0000ac1c8
	t.Logf("%T,%T", a, aPtr) //int,*int
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("空")
	}
	//输出 **  0  空
}
