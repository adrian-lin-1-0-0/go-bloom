package bloom

type Hasher interface {
	Hash(val []byte) int
}
