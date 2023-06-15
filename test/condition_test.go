package test

import (
	"testing"
)

func TestCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("not 0-3", i)
		}
	}
}
func TestCondition2(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even", i)
		case i%2 == 1:
			t.Log("Odd", i)
		default:
			t.Log("unknow", i)
		}
	}
}
