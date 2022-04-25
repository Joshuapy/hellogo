
package main


import "fmt"



func main() {
    var Add = func(a, b int) int {
        return a + b
    }
    fmt.Printf("type of Add: %T\n", Add)
    c := Add(1, 2)
    fmt.Printf("type of c: %T, and value: %v\n", c, c)
}
