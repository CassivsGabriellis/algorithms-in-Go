package main

import "fmt"

func binarySearch(array []int, num int) bool {
    //lo = lowest value
    lo := 0
    //hi = highest value
    hi := len(array)
    // while do not reach the highest value
    for lo < hi {
        // m = middle value
        m := (lo + hi) / 2
        // value that contains the middle as the index
        v := array[m]
        
        // Found the correct number (num)  
        if v == num {
            return true
                // if the value is greater than the number,
            // search the lower half of the array
        } else if v > num {
            hi = m
            // If the value is lower than the number,
            // search the upper half of the array
        } else {
            lo = m + 1
        }
    }

    return false
}

func main() {
    array := []int{1,2,3,4,5,6,7}
    number := 6
    found := binarySearch(array, number)
    fmt.Println(found)
}
