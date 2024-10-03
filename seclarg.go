package main

import (package main

    import (
    
    "fmt"
    
    "math"
    
    )
    
    func getSecLargest(vec []int32) int32 {
    
    fMax := int32(math.MinInt32)
    
    sMax := int32(math.MinInt32 + 1)
    
    for _, num := range vec {
    
    if num > fMax {
    
    fMax = num
    
    }
    
    }
    
    for _, num := range vec {
    
    if num < fMax {
    
    if num > sMax {
    
    sMax = num
    
    }
    
    }
    
    }
    
    return sMax
    
    }
    
    func main() {
    
    var n int32
    
    fmt.Scan(&n)
    
    vec := make([]int32, n)
    
    for i := range vec {
    
    fmt.Scan(&vec[i])
    
    }
    
    ans := getSecLargest(vec)
    
    fmt.Printf("\nSecond Largest : %d\n", ans)
    
    }
    "fmt"
    "math"
)

func getSecLargest(vec []int32) int32 {
    fMax := int32(math.MinInt32)
    sMax := int32(math.MinInt32 + 1)
    for _, num := range vec {
        if num > fMax {
            fMax = num
        }
    }
    for _, num := range vec {
        if num < fMax {
            if num > sMax {
                sMax = num
            }
        }
    }
    return sMax
}

func main() {
    var n int32
    fmt.Scan(&n)
    vec := make([]int32, n)
    for i := range vec {
        fmt.Scan(&vec[i])
    }
    ans := getSecLargest(vec)
    fmt.Printf("\nSecond Largest : %d\n", ans)
}