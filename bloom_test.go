package gobloom

import "testing"

func TestDefaultBloomSetAndExist(t *testing.T) {
	b := Default(30, 3)

	keys := []string{"hello", "world", "foo", "bar", "baz"}

	for _, key := range keys {
		b.Set(key)
	}

	for _, key := range keys {
		if !b.Exist(key) {
			t.Errorf("key %s should exist", key)
		}
	}
}

func TestDefaultBloomSetAndLen(t *testing.T) {
	b := Default(30, 3)

	keys := []string{"hello", "hello", "hello", "world", "world", "foo", "bar", "baz"}
	sets := make(map[string]struct{})
	for _, key := range keys {
		b.Set(key)
		sets[key] = struct{}{}
	}

	if len(sets) != b.Len() {
		t.Errorf("len should be %d, but got %d", len(sets), b.Len())
	}
}
