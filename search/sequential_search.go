package main

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