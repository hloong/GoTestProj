package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Huang", Age: 18}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "Roe"
	e2.Age = 28
	t.Log(e, e1, e1.Id, e2) //{0 Bob 20} { Huang 18}  &{2 Roe 28}
	t.Logf("e is %T", e)    //e is test.Employee
	t.Logf("e2 is %T", e2)  //e2 is *test.Employee
}

func TestStructOp(t *testing.T) {
	e := Employee{"0", "bod", 12}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log("string1--->", e.String1())
	t.Log("string2--->", e.String2())
	//Address is c0000a2490
	//Address is c0000a24c0    encap_test.go:30: string1---> ID:0/Name:bod/Age:12
	//Address is c0000a2490    encap_test.go:31: string2---> ID:0-Name:bod-Age:12
}

//这种定义方式在实例对应方法被调用时，实例的成员会进行值复制
func (e Employee) String1() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

//通常情况下为了避免内存拷贝，我们使用这种定义方式
func (e *Employee) String2() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
