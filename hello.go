package main

import "fmt"

func main(){
	var arr = [] int {1,2,3,4,5}
	fmt.Println(SumArr(arr))

	fmt.Println(SequentialSearch(arr, 6))

	fmt.Println(BinarySearch(arr, 1))
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
func BinarySearch(arr[] int, value int) bool{
	var j int
	low := 0
	high := len(arr)
	mid := low + (low + high)/2

	for j = 0; j < high; j++ {
		if arr[mid] == value {
			return true
		} else if value > arr[mid]{
			low = mid
		}else {
			high = mid + 1
		}
	}
	return false

}