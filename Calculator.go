/*
A simple canculator made in Golang.
This calculator calculates the Sum, Difference, Product, Quotient of two inputted numbers(float32).
*/

package main

import "fmt"

func calc(x, y float32)(sum, diff, product, div float32){
    sum = x + y
    
    if x<y{
        diff = y - x
    } else{
        diff = x - y
    }

    product = x * y

    if x<y{
        div = y / x
    } else{
        div = x / y
    }

    return 
 
}

func main() {
    var no1 float32
    fmt.Scanln(&no1)

    var no2 float32
    fmt.Scanln(&no2)

    sum, diff, product, div := calc(no1, no2)

    //fmt.Println(sum, diff, product, div)
    fmt.Println("The Sum is:        ", sum)
    fmt.Println("The Difference is: ", diff)
    fmt.Println("The Product is:    ", product)
    fmt.Println("The Quotient is:   ", div)
}
