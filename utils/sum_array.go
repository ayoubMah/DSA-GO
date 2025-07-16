package utils

func SumArr(arr[] int) int{
	var j int
	sum:= 0
	for j = 0; j < len(arr); j++{
		sum += arr[j]
	}
	return sum

}

