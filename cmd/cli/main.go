package main

import (
	"context"
	"fmt"
)

var (
	GlobalsAllTheTime string = ""
)

func FuncWithISSUEs(ctx context.Context) error {
	return nil
}

func main() {
	_ = FuncWithISSUEs(nil)

	fmt.Println("vim-go")
	fmt.Println("some other message!")
}
