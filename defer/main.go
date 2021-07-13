package main

import "fmt"

func main() {
	a()
	b()
	fmt.Println(c())
}

// 当defer被声明时，参数会实时解析，所以这里得到的结果是0
func a() {
	i := 0
	defer fmt.Println(i)
	i++
}

// golang按照先进后出的规则依次调用defer，所以这里的结果是3210
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// return 之后进行函数调用，这里先return 1，然后调用literal
func c() (i int) {
	defer func() { i++ }()
	return 1
}
