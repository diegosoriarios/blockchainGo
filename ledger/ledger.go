package ledger

import (
	"errors"
)

type Ledger struct {
	Genesis *Block
}

func (l *Ledger) Add(h string) string {
	if l.Genesis == nil {
		l.Genesis = &Block{Hash: h, Next: nil, Previous: nil}
		return h
	}

	block := l.Genesis

	for block.Next != nil {
		block = block.Next
	}

	block.Next = &Block{Hash: h, Next: nil, Previous: block}

	return h
}

func (l *Ledger) Get(h string) (*Block, error) {
	if l.Genesis == nil {
		return nil, errors.New("The chain is Empty")
	}

	block := l.Genesis
	hash := l.Genesis.Hash

	for block != nil {
		if hash == h {
			return block, nil
		}
		block = block.Next
		hash = block.Hash
	}

	return nil, errors.New("Block not found")

}
