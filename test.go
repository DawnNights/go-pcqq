package main

import (
	"fmt"
	"pcqq/utils"
)

func main() {
	bin := []byte("MzYzNDQxMzkzNQ==")
	fmt.Println(utils.Bin2HexTo(bin))
}
