package main

import (
	"pcqq/core"
)

func main() {
	pc := core.PCQQ{}
	pc.Init()
	pc.GetQrCode()
	// pc.LoadConfig()

	pc.ListenMessage()

}