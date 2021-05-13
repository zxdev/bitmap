package bitmap

import (
	"bufio"
	"io"
)

type BitMap []uint64

// NewBitMap populates and returns a *BitMap based on the io.Reader
// reading order and populates bitmap *map[string]int when non-nil; 
// expects an \n delimited simple reference list to be provided for 
// the bitmap lookup reference
func NewBitMap(r io.Reader, bitmap *map[string]int) *BitMap {

	var b = new(BitMap)

	if r != nil {
		var index int
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			b.Set(index)
			if bitmap != nil {
				(*bitmap)[scanner.Text()] = index
			}
			index++
		}
	}

	return b
}

// Set index; dynamic sizer
func (b *BitMap) Set(n int) {

	if n/64 < len(*b) {
		(*b)[n/64] |= 1 << (n % 64)
		return
	}

	var seg uint64
	*b = append(*b, seg)
	b.Set(n)
}

// Unset index
func (b *BitMap) Unset(n int) {

	if n/64 < len(*b) {
		(*b)[n/64] &^= 1 << (n % 64)
	}

}

// IsSet index
func (b *BitMap) IsSet(n int) bool {

	if n/64 < len(*b) {
		return (*b)[n/64]&(1<<(n%64)) != 0
	}
	return false

}
