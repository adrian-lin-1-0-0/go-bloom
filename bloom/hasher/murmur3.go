package hasher

import (
	"go-bloom/bloom"
	"math"

	"github.com/spaolacci/murmur3"
)

const is64Bit = uint64(^uintptr(0)) == ^uint64(0)

type Murmur3 struct {
}

var _ bloom.Hasher = (*Murmur3)(nil)

func NewMurmur3() bloom.Hasher {
	return &Murmur3{}
}

func (m *Murmur3) Hash(val []byte) int {
	if is64Bit {
		return m.hash64(val)
	}
	return m.hash32(val)
}

func (m *Murmur3) hash64(val []byte) int {
	hasher := murmur3.New64()
	_, _ = hasher.Write(val)
	return int(hasher.Sum64() % math.MaxInt64)
}

func (m *Murmur3) hash32(val []byte) int {
	hasher := murmur3.New32()
	_, _ = hasher.Write(val)
	return int(hasher.Sum32() % math.MaxInt32)
}
