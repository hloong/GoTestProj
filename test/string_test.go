package test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值""
	s = "Hello"
	t.Log(len(s)) //  5
	//s[1]='3'  //string是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据  严
	t.Log(s)           //严
	s = "中"
	t.Log(len(s)) //3
	c := []rune(s)
	t.Logf("中 unicode %x", c[0]) //中 unicode 4e2d
	t.Logf("中 UTF8 %x", s)       // 中 UTF8 e4b8ad
}

func TestStringToRune(t *testing.T) {
	s := "以太坊"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
	//以 20197
	//太 22826
	//坊 22346
}

func TestStringSlice(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part) //A B C
	}
	t.Log(strings.Join(parts, "-")) //A-B-C
}
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s) //str10
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i) //20
	}
}
