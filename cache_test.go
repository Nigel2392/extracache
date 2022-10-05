package main

import (
	"strconv"
	"testing"
)

func Test_Set(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Set("key"+strconv.Itoa(i), "value"+strconv.Itoa(i), 60)
	}
	for i := 0; i < 10; i++ {
		if cache.Data["key"+strconv.Itoa(i)].Value != "value"+strconv.Itoa(i) {
			t.Errorf("Cache does not have key %d", i)
		}
	}
	LOGGER.Test("Cache Test_Set Finished")
}
func Test_Get(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	for i := 0; i < 10; i++ {
		if cache.Get("key"+strconv.Itoa(i)) == nil {
			t.Errorf("Cache does not have key %d", i)
		}
	}
	cache.DeleteAll()
	if cache.Get("key") != nil {
		t.Errorf("Cache has key")
	}
	LOGGER.Test("Cache Test_Get Finished")
}
func Test_Delete(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	for i := 0; i < 10; i++ {
		cache.Delete("key" + strconv.Itoa(i))
	}
	if len(cache.Data) != 0 {
		t.Errorf("Cache is not empty")
	}
	LOGGER.Test("Cache Test_Delete Finished")
}
func Test_DeleteAll(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	cache.DeleteAll()
	if len(cache.Data) != 0 {
		t.Errorf("Cache is not empty")
	}
	LOGGER.Test("Cache Test_DeleteAll Finished")
}
func Test_GetSize(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	if cache.GetSize() != 10 {
		t.Errorf("Cache size is not 10")
	}

	LOGGER.Test("Cache Test_GetSize Finished")
}
func Test_GetKeys(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	keys := cache.GetKeys()
	if len(keys) != 10 {
		t.Errorf("Cache keys length is not 10")
	}
	LOGGER.Test("Cache Test_GetKeys Finished")
}
func Test_HasKey(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	for i := 0; i < 10; i++ {
		cache.Data["key"+strconv.Itoa(i)] = &CachedItem{Key: "key" + strconv.Itoa(i), Value: "value" + strconv.Itoa(i), TTL: 60}
	}
	for i := 0; i < 10; i++ {
		if !cache.HasKey("key" + strconv.Itoa(i)) {
			t.Errorf("Cache does not have key %d", i)
		}
	}
	LOGGER.Test("Cache Test_HasKey Finished")
}
func Test_ItemTTL(t *testing.T) {
	var cache Cache = Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	cache.Set("key", "value", 1)
	if cache.Data["key"].TTL != 1 {
		t.Errorf("Cache item TTL is not 1")
	}
	if cache.ItemTTL("key") != 1 {
		t.Errorf("Cache item TTL is not 1")
	}
	LOGGER.Test("Cache Test_ItemTTL Finished")
}
