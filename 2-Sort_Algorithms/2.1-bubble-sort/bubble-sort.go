package main

import "fmt"

func bubleSort(arr []int) {
    for i := 0; i < len(arr); i++ {
        //  The 'i' in the loop condition ensures that in each iteration of the outer loop, 
        // the largest elements are bubbled up to the end
        // of the array, so we don't need to compare them again.
        for j := 0; j < len(arr) - 1 - i; j++ {
            // if the current element arr[j] is greater than
            // the next element arr[j+1], a swap is performed
            // using a temporary variable 'temp'(-orary).
            if arr[j] > arr[j + 1] {
                temp := arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp
            }
        }
    }
}

func main() {
    arr := []int{2,6,5,1,0,9}
    bubleSort(arr)
    fmt.Println(arr)
}
