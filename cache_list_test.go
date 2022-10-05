package main

import (
	"testing"
)

func Test_GetCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.SetCache(i, &Cache{Channel_ID: i})
	}
	for i := 0; i < 10; i++ {
		for cache := range cacheList.Caches {
			flag := false
			for j := 0; j < 10; j++ {
				if cache == j {
					flag = true
				}
			}
			if !flag {
				t.Errorf("Cache %d is not in the cache list", cache)
			}
		}
	}
	LOGGER.Test("Cache_List Test_GetCache Finished")
}
func Test_SetCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.SetCache(i, &Cache{Channel_ID: i})
	}
	for i := 0; i < 10; i++ {
		for cache := range cacheList.Caches {
			flag := false
			for j := 0; j < 10; j++ {
				if cache == j {
					flag = true
				}
			}
			if !flag {
				t.Errorf("Cache %d is not in cache list", cache)
			}
		}
	}
	LOGGER.Test("Cache_List Test_SetCache Finished")
}
func Test_DeleteCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.SetCache(i, &Cache{Channel_ID: i})
	}
	for i := 0; i < 10; i++ {
		cacheList.DeleteCache(i)
	}
	for i := 0; i < 10; i++ {
		for cache := range cacheList.Caches {
			flag := false
			for j := 0; j < 10; j++ {
				if cache == j {
					flag = true
				}
			}
			if flag {
				t.Errorf("Cache %d is in cache list", cache)
			}
		}
	}
	LOGGER.Test("Cache_List Test_DeleteCache Finished")
}
func Test_SetItemInCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 0)
	}
	for cache := range cacheList.Caches {
		if !cacheList.Caches[cache].Data["key"].IsExpired() {
			t.Errorf("Cache %d is not expired", cache)
		}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 100)
	}
	for cache := range cacheList.Caches {
		if _, ok := cacheList.Caches[cache].Data["key"]; !ok {
			t.Errorf("Cache %d does not have key", cache)
		} else if cacheList.Caches[cache].Data["key"].IsExpired() {
			t.Errorf("Cache %d is expired", cache)
		}
	}
	LOGGER.Test("Cache_List Test_SetItemInCache Finished")
}
func Test_GetItemFromCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	for cache := range cacheList.Caches {
		if cacheList.GetItemFromCache(cache, "key") == nil {
			t.Errorf("Cache %d does not have key", cache)
		}
	}
	LOGGER.Test("Cache_List Test_GetItemFromCache Finished")
}
func Test_DeleteItemFromCache(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	for cache := range cacheList.Caches {
		cacheList.DeleteItemFromCache(cache, "key")
	}
	for cache := range cacheList.Caches {
		if _, ok := cacheList.Caches[cache].Data["key"]; ok {
			t.Errorf("Cache %d has key", cache)
		}
	}
	LOGGER.Test("Cache_List Test_DeleteItemFromCache Finished")
}
func Test_GetCacheSize(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	for cache := range cacheList.Caches {
		if cacheList.GetCacheSize(cache) != 1 {
			t.Errorf("Cache %d does not have size 1", cache)
		}
	}
	LOGGER.Test("Cache_List Test_GetCacheSize Finished")
}
func Test_CacheHasKey(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	for cache := range cacheList.Caches {
		if !cacheList.CacheHasKey(cache, "key") {
			t.Errorf("Cache %d does not have key", cache)
		}
	}
	LOGGER.Test("Cache_List Test_CacheHasKey Finished")
}
func Test_GetCacheKeys(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	for cache := range cacheList.Caches {
		keys := cacheList.GetCacheKeys(cache)
		if len(keys) != 1 {
			t.Errorf("Cache %d does not have size 1", cache)
		} else if keys[0] != "key" {
			t.Errorf("Cache %d does not have key", cache)
		}
	}
	LOGGER.Test("Cache_List Test_GetCacheKeys Finished")
}
func Test_GetCacheSizeAll(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	if cacheList.GetCacheSizeAll() != 10 {
		t.Errorf("Cache size is not 10")
	}
	LOGGER.Test("Cache_List Test_GetCacheSizeAll Finished")
}
func Test_GetCacheKeysAll(t *testing.T) {
	var cacheList CacheList = CacheList{Caches: make(map[int]*Cache)}
	for i := 0; i < 10; i++ {
		cacheList.Caches[i] = &Cache{Channel_ID: i, Data: make(map[string]*CachedItem)}
	}
	for cache := range cacheList.Caches {
		cacheList.SetItemInCache(cache, "key", "value", 60)
	}
	keys := cacheList.GetCacheKeysAll()
	if len(keys) != 10 {
		t.Errorf("Cache size is not 10")
	}
	LOGGER.Test("Cache_List Test_GetCacheKeysAll Finished")
}
