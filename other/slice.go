package other

import "fmt"

func Other(){
	s := [...]int{0: 1, 2: 3, 4: 9, 8: 10}
	fmt.Printf("s: %v\n", s)
	t := [...]int{1, 2, 3, 4: 9, 8: 10, 1, 12}
	fmt.Printf("t: %v\n", t)
	fmt.Printf("len(t): %d\n", len(t))
	fmt.Printf("type(t): %T\n", t)
	fmt.Println("================================")
	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界
	fmt.Println("================================")
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
	fmt.Println("================传统方法================")
	const o = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(o); i++ {
		fmt.Printf("%d %x\n", i, o[i])
	}
	fmt.Println("================slice copy() ================")
	array1 := [3]int{1, 2, 3}
	slice1 := array1[:]
	fmt.Println(slice1)
	slice1 = append(slice1, 0)
	fmt.Println(slice1)
	copy(slice1[1:], slice1[0:])
	fmt.Println(slice1)
	slice1[0] = 0
	fmt.Println(slice1)

	fmt.Println("================slice 中间插入数组 ================")
	fmt.Println("slice1:", slice1)
	x := []int{10, 12, 13}
	fmt.Println("x:", x)

	i := 2
	slice1 = append(slice1, x...) // 为x切片扩展足够的空间
	fmt.Println("slice1:", slice1)
	copy(slice1[i+len(x):], slice1[i:]) // a[i:]向后移动len(x)个位置
	fmt.Println("slice1:", slice1)
	copy(slice1[i:], x) // 复制新添加的切片
	fmt.Println("slice1:", slice1)

	fmt.Println("================slice 验证copy ================")
	slice2 := []int{0, 1, 2, 3}
	copy(slice2[1:], []int{5})
	fmt.Println(slice2)

	fmt.Println("================slice 验证copy ================")
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	fmt.Println(a)
	a = a[:copy(a, a[3:])] // 删除开头N个元素
	fmt.Println(a)

	fmt.Println("============== defer 匿名函数 ================")
	v := Inv()
	fmt.Println(v)

	fmt.Println("============== defer 匿名函数 ================")
	mm := int64(1)
	fmt.Println(float64(mm) / 100)

	fmt.Println("============== 方法 ================")
	boy := new(Boy)
	(*Boy).Say(boy, "1123")
}