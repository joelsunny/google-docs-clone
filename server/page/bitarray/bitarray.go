package bitarray

import "fmt"

var offsetArrOR = [8]int{0b1000_0000, 0b0100_0000, 0b0010_0000, 0b0001_0000, 0b0000_1000, 0b0000_0100, 0b0000_0010, 0b0000_0001}
var offsetArrAND = [8]int{0b0111_1111, 0b1011_1111, 0b1101_1111, 0b1110_1111, 0b1111_0111, 0b1111_1011, 0b1111_1101, 0b1111_1110}

// BitArray :- struct for bit array
type BitArray struct {
	Content []byte
	len     uint
}

// NewBitArray :- returns a bit array with specified length
func NewBitArray(len uint) *BitArray {
	b := &BitArray{Content: make([]byte, len>>3), len: (len >> 3) << 3}
	return b
}

// ToggleBit :- toggle specified bit
func (b *BitArray) ToggleBit(bitpos uint) {
	return
}

// SetBit :- set specified bit
func (b *BitArray) SetBit(bitpos uint) {
	if bitpos > b.len {
		fmt.Println("Error: bitpos out of range")
		return
	}

	bitposOffset := bitpos % 8
	b.Content[bitpos>>3] |= byte(offsetArrOR[bitposOffset])

}

// ResetBit :- reset specified bit
func (b *BitArray) ResetBit(bitpos uint) {
	if bitpos > b.len {
		fmt.Println("Error: bitpos out of range")
		return
	}

	bitposOffset := bitpos % 8
	b.Content[bitpos>>3] &= byte(offsetArrAND[bitposOffset])

}

// Print :- prints the bit array
func (b *BitArray) Print() {
	var str string
	for _byte := range b.Content {
		str += fmt.Sprintf("%08b ", b.Content[_byte])
	}
	fmt.Println(str)
}
