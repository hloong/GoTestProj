package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValus() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, _ := returnMultiValus()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 1
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2))    //1+1+2= 4
	t.Log(Sum(1, 2, 3)) //1+2+3=7
}

func Clear() {
	fmt.Println("Clear resources")
}
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("Error")
	//fmt.Println("End")
	//Start
	//End
	//Clear resources
}
