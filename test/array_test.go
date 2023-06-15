package test

import "testing"

func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5, 5, 6}
	arr1[1] = 5
	t.Log(arr1[1], arr[2]) //5,0
	t.Log(arr1, arr2)      //[1 5 3 4] [1 2 3 4 5 5 6]
}
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 5}
	for i := 0; i < len(arr3); i++ { //传统方法
		t.Log(arr3[i]) //1,3,5
	}
	for idx, e := range arr3 { //go提供了range可以替换上面的方法
		t.Log(idx, e) // idx属于索引，e数组里面的数据; 0 1,1 3,2 5,
	}
	for _, e := range arr3 { //不用索引，用_占位
		t.Log(e) //1,3,5
	}
}
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[:2]
	var arr3Sec2 = arr3[1:len(arr3)]
	arr3Sec3 := arr3[1:]
	t.Log(arr3Sec, "*", arr3Sec2, "*", arr3Sec3) //[1 2] * [2 3 4 5] * [2 3 4 5]
}
