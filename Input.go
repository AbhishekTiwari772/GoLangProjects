package main

import (
   "fmt"
   "bufio"
   "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    fmt.Println(input)
}
