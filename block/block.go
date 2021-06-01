package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Block
type Block struct {
	Idx          int
	Data         map[string]string
	Created      time.Time
	Hash         string
	PreviousHash string
	Nonce        int
	Next         *Block `json:"-"`
}

func (b *Block) MineBlock(difficulty int) {
	s := ""
	for i := 0; i < difficulty; i++ {
		s = s + "0"
	}
	for {
		b.Nonce += 1
		b.Hash = b.GenerateHash()
		if b.Hash[0:difficulty] == s {
			break
		}
	}
}

// NewBlock
func NewBlock(data map[string]string) *Block {
	b := &Block{
		Data:    data,
		Created: time.Now(),
		Nonce:   0,
	}
	return b
}

// GenerateHash
func (b *Block) GenerateHash() string {
	sha := sha256.New()
	str := fmt.Sprintf("%d%d%s%s%s",
		b.Nonce, b.Idx, b.PreviousHash, b.Created, b.DataToJSON())
	_, err := sha.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func (b *Block) Verify() bool {
	if b.Hash != b.GenerateHash() {
		return false
	}
	if b.Next == nil {
		return true
	}
	return b.Hash == b.Next.PreviousHash
}

func (b *Block) DataToJSON() string {
	bytes, err := json.Marshal(b.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func (b *Block) ToJSON() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
