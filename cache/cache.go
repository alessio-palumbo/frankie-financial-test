package cache

import "sync"

// SessionCache keeps track of the sessionKeys used thorughout the life
// of the running process to make sure they are unique.
type SessionCache struct {
	sync.RWMutex
	sessionKeys map[string]bool
}

// New initialise a SessionCache
func New() *SessionCache {
	return &SessionCache{
		sessionKeys: make(map[string]bool),
	}
}

// Has checks whether a session key has already been used
func (sc *SessionCache) Has(key string) bool {
	sc.RLock()
	_, found := sc.sessionKeys[key]
	sc.RUnlock()

	return found
}

// Store saves the session key in the cache
func (sc *SessionCache) Store(key string) {
	sc.Lock()
	sc.sessionKeys[key] = true
	sc.Unlock()
}
