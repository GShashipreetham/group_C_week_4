package main

func BubbleSort(arr []int) {
	s := len(arr)
	for i := 0; i < s-1; i++ {
		for j := 0; j < s-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
