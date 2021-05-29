/*
This GO code takes input as an integer and prints all the whole numbers from 0 to the inputted integer.
*/

package main

import "fmt"

func main() {
    var input int
    fmt.Scanln(&input)
    
    for i := 0; i <= input; i++{
        fmt.Println(i)
    }
}
