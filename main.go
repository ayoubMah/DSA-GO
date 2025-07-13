package main

import(
	"fmt"
  
)

func main(){
	var arr = [] int {1,2,3,4,5}

	fmt.Println(SumArr(arr))

	fmt.Println(SequentialSearch(arr, 6))

	fmt.Println(BinarySearch(arr, 33))
}