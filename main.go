package main

import(
	"fmt"
	"dsa-go/search"
	"dsa-go/utils"
)

func main(){
	var arr = [] int { 1, -2, 3, 4, -4, 6, -14, 6, 2 }

	fmt.Println(utils.SumArr(arr))

	fmt.Println(search.SequentialSearch(arr, 6))

	fmt.Println(search.BinarySearch(arr, 3))

	var currentArray [] int
	for x:= 0; x < len(arr); x++{
		c
		fmt.Println(arr)
	}

	var i int
	var myArr [] int
	sum:= 0
	for i = 0 ; i < len(arr); i++{
		//fmt.Println(arr[i])
		sum += arr[i]
		//fmt.Println(sum)
		myArr = append(myArr,sum)
		fmt.Println(myArr)
		fmt.Println("=================")
	}
	max := myArr[0]
	for j:= 0; j < len(myArr) ; j++{
		if myArr[j] > max{
			max = myArr[j]
		}
	}
	fmt.Println(max)
	fmt.Println("=================")



}