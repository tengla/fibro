package main

import (
	"flag"
	"fmt"

	"github.com/tengla/fibro/block"
)

var difficulty = flag.Int("difficulty", 2, "The mining difficulty")

func main() {

	flag.Parse()

	chain := block.CreateChain(*difficulty)

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the first",
	}))

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the second",
	}))

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the third", "Note": "Funny",
	}))

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the fourth",
	}))

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the fifth",
	}))

	chain.AddBlock(block.NewBlock(map[string]string{
		"Name": "I am the sixth",
	}))

	chain.EveryBlock(func(b *block.Block) {
		fmt.Printf("%s\n", b.ToJSON())
	})
}
