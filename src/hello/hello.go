package main

import ( 
    "fmt"
    "math"
    "math/rand"

)

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The square root of 16 is", math.Sqrt(16))
    square()
    random()
}

func square() {
    fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func random() {
    fmt.Println("Your random number is:", rand.Intn(100))
}
