package main

import (
    "fmt"
    "math"
)

func twoCrystalBalls(breaks []bool) int { //return the index of the array
    
    // jumping by the square of n in the arrayi
    // It is, the exactly amount to jump 
    jumpAmount := int(math.Sqrt(float64(len(breaks))))
    
    // 1. Using the first crystal ball to see where does it break
    // 2. Trying until the ball breaks  
    // 3. Find the point where it breaks

    // Using the first crystal ball to find the potential
    // range where it breaks
	i := jumpAmount
	for i < len(breaks) && !breaks[i] {
		i += jumpAmount
	}
               
    //4. Return to the beginning of the array
    //5. Walk foward in the array
    //6. walk the maximum of jump amount of steps

    // Linear search within the potential range
    // to find the exact breaking point
	for j := i - jumpAmount + 1; j <= i && i < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}

func main() {
    breaks := []bool{false, false, true, false, false, true, false, false, false, false, false, true}
    result := twoCrystalBalls(breaks)
    fmt.Println(result)
}
