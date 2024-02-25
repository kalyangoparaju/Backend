package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("GOARCH:", os.Getenv("GOARCH"))
    fmt.Println("GOOS:", os.Getenv("GOOS"))
}
