package main

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