package engine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/emirpasic/gods/maps/hashmap"
)

type Engine struct {
	data  *hashmap.Map
	mutex sync.Mutex
}

func NewEngine() *Engine {
	return &Engine{
		data: hashmap.New(),
	}
}

func (e *Engine) Put(key string, value []byte) error {
	// validate the key. in the case that the key is too long we can't allow that for
	// performance issues, so we return an error before even trying to put the key
	// in the store.
	if !e.ValidateKey(key) {
		return errors.New("couldn't validate key")
	}

	// then, try to acquire the lock to insert the value in the store, replacing any
	// existent values with the same key.
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.data.Put(key, value)

	return nil
}

func (e *Engine) Get(key string) ([]byte, error) {
	// we're trading the performance overhead of if statements for the opportunity
	// to do as few lookups to the store as possible, so we want to make sure that,
	// before we do an actual lookup, we know that there actually *is* a posibility
	// for the key to be in the store.
	if !e.ValidateKey(key) {
		return nil, fmt.Errorf("key \"%s\" doesn't exist in the store", key)
	}

	// try to gain the lock of the mutex so we can grab the value. i really don't
	// know if we need this on get, though.
	e.mutex.Lock()
	defer e.mutex.Unlock()

	value, found := e.data.Get(key)

	if !found {
		// TODO: DRY
		return nil, fmt.Errorf("key \"%s\" doesn't exist in the store", key)
	}

	// make sure that it's a []byte underneath
	switch v := value.(type) {
	case []byte:
		return v, nil
	default:
		return nil, fmt.Errorf("unexpected value type for key %s", key)
	}
}

func (e *Engine) ValidateKey(key string) bool {
	return ValidateKey(key)
}
