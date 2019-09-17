package bitarray

import "testing"

func TestSetBit(t *testing.T) {
	barray := NewBitArray(8)
	barray.SetBit(0)
	if barray.Content[0] != byte(0b1000_0000) {
		t.Errorf("set bit failed")
	}
}

func BenchmarkSetBit(b *testing.B) {
	barray := NewBitArray(1024)
	for i := 0; i < b.N; i++ {
		barray.SetBit(1023)
	}
}

func BenchmarkNewBitArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewBitArray(128)
	}
}
