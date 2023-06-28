package main

import "fmt"

func linearSearch(list []int, number int) bool {
    for _, i := range list {
        if i == number {
            return true
        }
    }
    return false
}

func main() {
    palheiro := []int{1,2,3,4,5,6}
    agulha := 7
    achado := linearSearch(palheiro, agulha)
    fmt.Println(achado)
}
