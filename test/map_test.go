package map_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	t.Log(m["one"])              //1
	t.Logf("len m = %d", len(m)) //len m = 2
	m1 := map[string]int{}
	m1["one"] = 1
	t.Logf("len m1=%d", len(m1)) //len m1=1
	m2 := make(map[string]int, 10)
	t.Logf("len m2=%d", len(m2)) //len m2=0
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) //go不存在空指针，会默认给0
	m1[2] = 0
	t.Log(m1[2])
	//m1[3]= 0
	if v, ok := m1[3]; ok {
		t.Logf("key3 is %d", v) //如果赋值就存在，不赋值就不存在，不会报错
	} else {
		t.Log("key3 is not ")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 3, 3: 9}
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestMapFactory(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](3), m[2](3), m[3](3)) //  返回 3 9 27
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	printSet(n, mySet, t)
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	printSet(n, mySet, t)
	n = 3
	printSet(n, mySet, t)
}
func printSet(n int, set map[int]bool, t *testing.T) {
	if set[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
