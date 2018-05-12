package main

import ( 
    "fmt"
    "math"
    "math/rand"

)
var num1, num2 float64 = 7.3, 2.6 // declare variable outside func
var num6, num7 float32 = 4.5, 3.2

const integer int = 7
const text string = "hello"

func main() {
    fmt.Println("hello, world\n")
    fmt.Println("The square root of 16 is", math.Sqrt(16))
    square()
    random()
    // num3, num4 := 3.7, 9.8 // declare var inside func - is guessing status as float64?, no need to specify
    fmt.Println(add(num6, num7))
    fmt.Println(integer, text)
    fmt.Println(twoStrings("hi", "there"))
}

func square() {
    fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func random() {
    fmt.Println("Your random number is:", rand.Intn(1000))
}

func add(x, y float32) float32 {
    return x + y
}

func twoStrings(a, b string) (string, string) { // have to specify both types is returned separately
    return a, b
}
