package gobloom

import (
	"go-bloom/bloom"
	"go-bloom/bloom/hasher"
)

func New(bitmapLen, hashFuncLen int, h bloom.Hasher) *bloom.Bloom {
	return bloom.NewBloom(bitmapLen, hashFuncLen, h)
}

func Default(bitmapLen, hashFuncLen int) *bloom.Bloom {
	return bloom.NewBloom(bitmapLen, hashFuncLen, hasher.NewMurmur3())
}
