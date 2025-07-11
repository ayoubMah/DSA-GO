package main

import "fmt"

func main(){
	var arr = [] int {1,2,3,4,5}
	fmt.Println(SumArr(arr))

	fmt.Println(SequentialSearch(arr, 6))

	fmt.Println(BinarySearch(arr, 33))
}


func SumArr(arr[] int) int{
	var j int
	sum:= 0
	for j = 0; j < len(arr); j++{
		sum += arr[j]
	}
	return sum

}

func SequentialSearch(arr[] int, value int) bool {

	var j int
	for j = 0; j < len(arr) ; j++{
		if arr[j] == value{
			break
		}
	}
	if j == len(arr) {
		return false
	}
	return true

}

// TODO => NOT COMPLETED
func BinarySearch(arr[] int, target int) bool{

	low := 0
	high := len(arr)

	for low < high{
		mid := (int)(low + (high -  low)/2)
		value := arr[mid]
		if value == target {
			return true
		}else if value > target{
			high = mid
		}else {low = mid + 1 }
	}
	return false
}