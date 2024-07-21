package bloom

import (
	"strconv"
)

const is64Bit = uint64(^uintptr(0)) == ^uint64(0)

type Bloom struct {
	bitmapLen   int
	hashFuncLen int
	n           int
	bitmap      []int
	hasher      Hasher
	arch        int
}

func NewBloom(bitmapLen, hashFuncLen int, h Hasher) *Bloom {

	if h == nil {
		panic("hasher is nil")
	}

	arch := 32
	if is64Bit {
		arch = 64
	}

	return &Bloom{
		bitmapLen:   bitmapLen,
		hashFuncLen: hashFuncLen,
		bitmap:      make([]int, bitmapLen/arch+1),
		hasher:      h,
		arch:        arch,
	}
}

func (b *Bloom) Set(key string) {
	b.n++
	for _, offset := range b.getHashs(key) {
		index := offset / b.arch
		bitOffset := offset % b.arch
		b.bitmap[index] |= (1 << bitOffset)
	}
}

func (b *Bloom) Exist(key string) bool {
	for _, offset := range b.getHashs(key) {
		index := offset / b.arch
		bitOffset := offset % b.arch

		if b.bitmap[index]&(1<<bitOffset) == 0 {
			return false
		}
	}

	return true
}

func (b *Bloom) getHashs(val string) []int {
	hasheds := make([]int, 0, b.hashFuncLen)
	for i := 0; i < b.hashFuncLen; i++ {
		hashed := b.hasher.Hash([]byte(val))
		hasheds = append(hasheds, hashed%b.bitmapLen)
		if i == b.hashFuncLen-1 {
			break
		}
		val = strconv.FormatInt(int64(hashed), 10)
	}
	return hasheds
}
