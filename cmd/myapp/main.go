package main

import (
    "fmt"
    "my-go-project/internal/config"
)

func main() {
    cfg := config.LoadConfig()
    fmt.Println("App running with config:", cfg)
}
