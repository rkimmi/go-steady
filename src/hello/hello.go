package main

import ( 
    "fmt"
    "math"
    "math/rand"

)
var num1 float64 = 7.3
var num2 float64 = 2.6

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The square root of 16 is", math.Sqrt(16))
    square()
    random()
    fmt.Println(add(num1, num2))
}

func square() {
    fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func random() {
    fmt.Println("Your random number is:", rand.Intn(1000))
}

func add(x float64, y float64) float64 {
    return x + y
}
