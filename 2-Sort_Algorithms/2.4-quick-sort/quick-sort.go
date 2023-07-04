package main

import "fmt"

//Function to separate the partitions of the inputs
func partition(arr []int, low int, high int) int {
	pivot := arr[high] // Pivot as the last element in the array

	index := low - 1 // index of the smaller element

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			index++
			arr[i], arr[index] = arr[index], arr[i] // swap the elements
		}
	}

	index++
	arr[high], arr[index] = arr[index], pivot // place the pivot in the correct position

	return index //return the index of the pivot
}

//Quick Sort function
func qs(arr []int, low int, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)

	qs(arr, low, pivotIndex-1)  // recursively sort the left part of the array
	qs(arr, pivotIndex+1, high) // recursively sort the right part of the array
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	QuickSort(array)
	fmt.Println(array)
}