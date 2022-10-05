package main

type CacheList struct {
	Caches map[int]*Cache
}

// Get a cache
func (c *CacheList) GetCache(channel int) *Cache {
	MUTEX.Lock()
	cache := c.Caches[channel]
	if cache == nil {
		cache = &Cache{Channel_ID: channel}
		c.Caches[channel] = cache
	}
	MUTEX.Unlock()
	return cache
}

// Set a new cache
func (c *CacheList) SetCache(channel int, cache *Cache) {
	MUTEX.Lock()
	c.Caches[channel] = cache
	MUTEX.Unlock()
}

// Delete a cache
func (c *CacheList) DeleteCache(channel int) {
	delete(c.Caches, channel)
}

// Set item in a cache
func (c *CacheList) SetItemInCache(channel int, key string, value interface{}, ttl int) {
	c.GetCache(channel).Set(key, value, ttl)
}

// Get item from a cache
func (c *CacheList) GetItemFromCache(channel int, key string) interface{} {
	return c.GetCache(channel).Get(key)
}

// Delete a key from a cache
func (c *CacheList) DeleteItemFromCache(channel int, key string) bool {
	return c.GetCache(channel).Delete(key)
}

// Get size of a cache
func (c *CacheList) GetCacheSize(channel int) int {
	return c.GetCache(channel).GetSize()
}

// Verify if a key exists in a cache
func (c *CacheList) CacheHasKey(channel int, key string) bool {
	return c.GetCache(channel).HasKey(key)
}

// Get all keys from a cache
func (c *CacheList) GetCacheKeys(channel int) []string {
	return c.GetCache(channel).GetKeys()
}

// Get size of all caches
func (c *CacheList) GetCacheSizeAll() int {
	var size int
	for _, cache := range c.Caches {
		size += cache.GetSize()
	}
	return size
}

// Get all keys from all caches
func (c *CacheList) GetCacheKeysAll() []string {
	var keys []string
	for _, cache := range c.Caches {
		keys = append(keys, cache.GetKeys()...)
	}
	return keys
}
