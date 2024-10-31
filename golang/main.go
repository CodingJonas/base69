package main

import (
	"fmt"

	"github.com/pshihn/base69/base69"
	"github.com/pshihn/base69/cli"
)

func main() {
	config := cli.GetConfig()
	result := ""
	if config.DoDecode {
		result = base69.Decode([]rune(config.Value))
	} else {
		result = base69.Encode([]byte(config.Value))
	}
	fmt.Println(result)
}
