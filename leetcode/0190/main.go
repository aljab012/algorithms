package main

func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		ret = ret<<1 + num&1
		num >>= 1
	}
	return ret
}
