// testsimplesyncmap
package syncstrmap

import (
	"testing"
)

func TestNewSimpeSyncMap(t *testing.T) {
	m := NewSimpeSyncMap()
	if m == nil {
		t.Error("NewSimpeSyncMap: map is nil")
	}
}

func TestSet(t *testing.T) {
	m := NewSimpeSyncMap()
	if m == nil {
		t.Error("NewSimpeSyncMap: map is nil")
	}
	m.Set("One", 1)
	m.Set("Two", 2)
	if m.Size() != 2 {
		t.Error("syncmap length should be 2")
	}
}

func TestHas(t *testing.T) {
	m := NewSimpeSyncMap()
	if m == nil {
		t.Error("NewSimpeSyncMap: map is nil")
	}
	m.Set("One", 1)
	m.Set("Two", 2)

	if !m.Has("One") {
		t.Error("NewSimpeSyncMap: One should be in map")
	}
	if m.Has("Three") {
		t.Error("NewSimpeSyncMap: Three should be not in map")
	}
}

func TestGet(t *testing.T) {
	m := NewSimpeSyncMap()
	if m == nil {
		t.Error("NewSimpeSyncMap: map is nil")
	}
	m.Set("One", 1)
	m.Set("Two", 2)

	if val, ok := m.Get("One"); !ok {
		t.Error("NewSimpeSyncMap: One should be in map")
	} else {
		if val.(int) != 1 {
			t.Error("NewSimpeSyncMap: Key \"One\" value should be 1")
		}
	}
}
