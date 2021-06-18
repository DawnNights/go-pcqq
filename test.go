package main

import (
	"fmt"
	"pcqq/utils"
	"time"
)

func main() {
	fmt.Println(utils.Int64ToBytes(time.Now().Unix())[4:])
	fmt.Println(utils.Flip(utils.Int64ToBytes(time.Now().Unix())[4:]))
}