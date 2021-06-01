package block

import (
	"time"
)

// Chain
type Chain struct {
	Created        time.Time
	Genesis        *Block
	NumberOfBlocks int
	Difficulty     int
}

// IncrNumberOfBlocks
func (c *Chain) IncrNumberOfBlocks() int {
	c.NumberOfBlocks = c.NumberOfBlocks + 1
	return c.NumberOfBlocks
}

// CreateChain
func CreateChain(difficulty int) *Chain {
	chain := &Chain{
		Created:        time.Now(),
		NumberOfBlocks: 0,
		Difficulty:     difficulty,
	}
	b := NewBlock(map[string]string{
		"Name": "Genesis Block",
	})
	b.PreviousHash = "0000"
	b.Hash = b.GenerateHash()
	chain.Genesis = b
	return chain
}

// AddBlock
func (c *Chain) AddBlock(b *Block) {
	block := c.Genesis
	prevHash := block.Hash
	for block.Next != nil {
		block = block.Next
		prevHash = block.Hash
	}
	b.PreviousHash = prevHash
	b.Idx = c.IncrNumberOfBlocks()
	b.MineBlock(c.Difficulty)
	block.Next = b
}

// Validate
func (c *Chain) Validate() bool {
	v := true
	c.EveryBlock(func(b *Block) {
		if v {
			v = b.Verify()
		}
	})
	return v
}

// EveryBlock
func (c *Chain) EveryBlock(fn func(b *Block)) {
	cur := c.Genesis
	for cur.Next != nil {
		fn(cur)
		cur = cur.Next
	}
	fn(cur)
}
