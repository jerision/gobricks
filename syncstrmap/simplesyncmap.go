// simplesynmap:thread safe map. go orignal map with lock, key mast be type of string
package syncstrmap

import (
	"sync"
)

type SimpleSyncMap struct {
	items map[string]interface{}
	sync.RWMutex
}

//simple sync map
func NewSimpeSyncMap() *SimpleSyncMap {
	m := new(SimpleSyncMap)
	m.Lock()
	defer m.Unlock()
	m.items = make(map[string]interface{})
	return m
}

// find the value of key in the map
func (m SimpleSyncMap) Get(key string) (val interface{}, ok bool) {
	m.RLock()
	defer m.RUnlock()
	val, ok = m.items[key]
	return
}

// judge if the key is in the map
func (m SimpleSyncMap) Has(key string) (ok bool) {
	m.RLock()
	defer m.RUnlock()
	_, ok = m.items[key]
	return
}

//insert pair key,val into the map
func (m *SimpleSyncMap) Set(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[key] = val
}

//delete a item of key from map
func (m *SimpleSyncMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, key)
}

// get the size of map
func (m *SimpleSyncMap) Size() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.items)
}

// flush the map to make a new one
func (m *SimpleSyncMap) Flush() {
	m.Lock()
	defer m.Unlock()
	m.items = make(map[string]interface{})
}

// Item is a pair of key and value
type Item struct {
	Key   string
	Value interface{}
}

// Return a channel from which each item (key:value pair) in the map can be read
func (m *SimpleSyncMap) IterItems() <-chan Item {
	ch := make(chan Item)
	go func() {
		m.RLock()
		for key, value := range m.items {
			ch <- Item{key, value}
		}
		m.RUnlock()
		close(ch)
	}()
	return ch
}
