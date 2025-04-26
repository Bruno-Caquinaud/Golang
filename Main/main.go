package main

import (
    "fmt"

    "example.com/mypackage"
)

func main() {
    // Get a greeting message and print it.
    message := mypackage.Hello("Gladys")
    fmt.Println(message)
}
