package cli

import (
	"flag"
	"fmt"

	"github.com/pshihn/base69/log"
)

type Config struct {
	Value    string
	DoDecode bool
}

func NewConfig(value string, doDecode bool) *Config {
	return &Config{Value: value, DoDecode: doDecode}
}

func GetConfig() *Config {
	doDecode := flag.Bool("d", false, "decode input")
	flag.Usage = func() {
		fmt.Println("Usage: base69 [OPTION]... VALUE\n\nOptions:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) > 1 {
		log.ExitWithError(fmt.Sprintf("only one input string allowed, but %v provided", flag.Args()))
	}
	value := flag.Arg(0)

	return NewConfig(value, *doDecode)
}

func (c *Config) Print() {
	fmt.Println("decode?", c.DoDecode)
	fmt.Println("value to decode:", c.Value)
}
