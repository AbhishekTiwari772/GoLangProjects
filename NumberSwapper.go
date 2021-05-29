package main
import "fmt"

func swap(x, y int)(int, int){
    return y, x
}

func main() {
    var input1 int
    var input2 int

    fmt.Println("Enter 1st number: ")
    fmt.Scanln(&input1)
    fmt.Println("Enter 2nd number: ")
    fmt.Scanln(&input2)

    fmt.Println(swap(input1, input2))
}
