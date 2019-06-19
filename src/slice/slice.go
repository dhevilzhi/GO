package main

import "fmt"

func displaySlice(mySlice []int) {
	fmt.Println("len(mySlice):", len(mySlice)) // 数组元素个数
	fmt.Println("cap(mySlice):", cap(mySlice)) // 数组存储数量

	fmt.Print("Elements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println()
}

func main() {
	/*
		// 定义数组
		var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		// 打印数组信息
		fmt.Print("Elements of myArray: ")
		for _, v := range myArray {
			fmt.Print(v, " ")
		}
		fmt.Println()
	*/

	// 基于数组创建一个数组切片
	//	var mySlice []int = myArray[:5] // 前5个元素
	//	var mySlice []int = myArray[5:]		// 后5个元素
	//	var mySlice []int = myArray[:]		// 所有元素

	// 直接创建数组切片
	//	var mySlice = make([]int, 5) // 元素个数为5的数组切片, 元素初始值为0
	var mySlice = make([]int, 5, 10) // 元素个数为5的数组切片, 元素初始值为0, 并预留10个元素的存储空间
	//	var mySlice = []int {1, 2, 3, 4, 5}
	displaySlice(mySlice)

	// 数组切片追加元素
	//	mySlice = append(mySlice, 1, 2, 3, 4)	// 直接追加元素
	mySlice2 := []int{5, 6, 7, 8}
	mySlice = append(mySlice, mySlice2...) // mySlice2后边的需要加上"...", 否则编译出错
	displaySlice(mySlice)

	// 基于数组切片创建新的数组切片
	mySlice3 := mySlice[5:]
	displaySlice(mySlice3)

	// copy数组切片
	copySlice1 := []int{1, 2, 3, 4, 5}
	copySlice2 := []int{5, 4, 3}
	//	copy(copySlice2, copySlice1)	// 复制copySlice1的前3个元素到copySlice2中
	copy(copySlice1, copySlice2) // 复制copySlice2的元素到copySlice1中的前3个位置
	displaySlice(copySlice1)
	displaySlice(copySlice2)
}
