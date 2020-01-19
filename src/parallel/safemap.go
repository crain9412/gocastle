package parallel

import (
	"sync"
)

/*
	Thread-safe map of string -> []string, useful for forward lookup in a search engine for example
*/

type Value struct {
	words []string
}

type SafeMap struct {
	value map[string]*Value
	mux   sync.Mutex
}

func InitSafeMap() SafeMap {
	return SafeMap{value: make(map[string]*Value)}
}

func (safeMap *SafeMap) Put(key string, words []string) {
	safeMap.mux.Lock()
	value := &Value{words: words}
	safeMap.value[key] = value
	safeMap.mux.Unlock()
}

// Return empty list of strings and false if key isn't found, list of strings and true if key is found
func (safeMap *SafeMap) Get(key string) ([]string, bool) {
	safeMap.mux.Lock()
	defer safeMap.mux.Unlock()
	safeMapValue := safeMap.value[key]
	if safeMapValue == nil {
		return []string{}, false
	} else {
		return safeMapValue.words, true
	}
}
