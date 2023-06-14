package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) //0 0
	s0 = append(s0, 1)      //append 函数填充元素
	t.Log(len(s0), cap(s0)) //1 1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) // 4 4

	s2 := make([]int, 3, 5)    //make函数，表示创建一个长度为3，数量为5的int数组
	t.Log(len(s2), cap(s2))    // 3 5
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0
	s2 = append(s2, 1, 3, 4)   // 给函数填充,超过了原来的5，那么cap(s2)=5*2,每次不够用的时候就*2
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2)) //6 10
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := month[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) //[Apr May Jun] 3 9   cap输出9是因为截断数据从Apr开始，后面有9个数据
	summer := month[5:8]
	t.Log(summer, len(summer), cap(summer)) //[Jun Jul Aug] 3 7  cap输出为7，截断数据从Jun开始后面有7个数据
	summer[0] = "U"                         //summer赋值会影响Q2和month，因为他们是共享空间
	t.Log(summer, Q2)                       //[U Jul Aug] [Apr May U]
	t.Log(month)                            //[Jan Feb Mar Apr May U Jul Aug Sep Oct Nov Dec]
}

func TestSliceArrayCompare(t *testing.T) {
	c := [...]int{1, 2, 3}
	d := [...]int{1, 2, 3}
	if c == d {
		t.Log("c==d")
	}
	//a := []int{1, 2, 3}
	//b := []int{1, 2, 3}
	//if a == b { //会报错 invalid operation: a == b (slice can only be compared to nil)
	//	t.Log("a=b")
	//} else {
	//	t.Log("a!=b")
	//}
}
