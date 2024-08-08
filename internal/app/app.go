package main

import (
	"fmt"
	"messenger/internal/config"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
}
