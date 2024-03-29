package example

import "sync"

// Code generated by mapgen (https://github.com/s-kirby/mapgen), DO NOT EDIT.

// ChannelIntMap is a generated thread safe map
// with key Channel and value int
type ChannelIntMap struct {
	sync.Mutex

	// M contains the underlying map.
	// Goroutines which access M directly should hold
	// the mutex.
	M map[Channel]int
}

// NewChannelIntMap returns an instantiated thread safe map
// with key Channel and value int
func NewChannelIntMap() *ChannelIntMap {
	return &ChannelIntMap{
		M: make(map[Channel]int),
	}
}

// Open allows a closure to safely operate on the map
func (m *ChannelIntMap) Open(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// Copy generates a copy of the map.
func (m *ChannelIntMap) Copy() (c map[Channel]int) {
	c = make(map[Channel]int, len(m.M))

	m.Lock()

	for k, v := range m.M {
		c[k] = v
	}

	m.Unlock()

	return c
}

// Set sets a key on the map
func (m *ChannelIntMap) Set(key Channel, val int) {
	m.Lock()
	m.M[key] = val
	m.Unlock()
}

// SetIfNotExist sets a key on the map if it doesn't exist.
// It returns the value which is set.
func (m *ChannelIntMap) SetIfNotExist(key Channel, val int) int {
	m.Lock()
	v, ok := m.M[key]
	if !ok {
		m.M[key] = val
	} else {
		val = v
	}
	m.Unlock()
	return val
}

// Delete removes a key from the map
func (m *ChannelIntMap) Delete(key Channel) {
	m.Lock()
	delete(m.M, key)
	m.Unlock()
}

// Get retrieves a key from the map
func (m *ChannelIntMap) Get(key Channel) int {

	m.Lock()
	v := m.M[key]
	m.Unlock()

	return v
}

// Len returns the length of the map
func (m *ChannelIntMap) Len() int {

	m.Lock()
	n := len(m.M)
	m.Unlock()

	return n
}

// GetEx retrieves a key from the map
// and whether it exists
func (m *ChannelIntMap) GetEx(key Channel) (int, bool) {

	m.Lock()
	v, exists := m.M[key]
	m.Unlock()

	return v, exists
}

// Exists returns if a key exists
func (m *ChannelIntMap) Exists(key Channel) bool {

	m.Lock()
	_, exists := m.M[key]
	m.Unlock()

	return exists
}
