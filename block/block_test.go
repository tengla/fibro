package block

import (
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	b := NewBlock(map[string]string{
		"Name": "Test",
	})
	if b == nil {
		t.Fatal("Could not create Block")
	}
	if b.Data["Name"] != "Test" {
		t.Fatal("Block.Data['Name'] is not 'Test'")
	}
	// should be very recent
	if time.Since(b.Created) > time.Duration(time.Second) {
		t.Fatal("Block.Created was not set correctly")
	}
}

func TestMineBlock(t *testing.T) {
	// g for genesis
	g := NewBlock(map[string]string{
		"Name": "Genesis",
	})
	g.Idx = 1
	g.PreviousHash = "0000"
	g.Hash = g.GenerateHash()
	if g.Hash != g.GenerateHash() {
		t.Fatal("Hash is not deterministic.")
	}
}
