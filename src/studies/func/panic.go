package main

import "fmt"

func main() {
    defer func() {
        x := recover()
        fmt.Println(x)
    }()
    panic("Pânico")
}