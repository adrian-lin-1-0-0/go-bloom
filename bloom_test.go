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
