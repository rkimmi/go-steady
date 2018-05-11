package main

import ( 
    "fmt"
    "math"
    "math/rand"

)
var num1, num2 float64 = 7.3, 2.6 // declare variable outside func

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The square root of 16 is", math.Sqrt(16))
    square()
    random()
    num3, num4 := 3.7, 9.8 // declare var inside func - is guessing status as float, no need to specify
    fmt.Println(add(num1, num2))
    fmt.Println(add(num3, num4))
}

func square() {
    fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func random() {
    fmt.Println("Your random number is:", rand.Intn(1000))
}

func add(x, y float64) float64 {
    return x + y
}
