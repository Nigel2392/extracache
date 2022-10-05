package main

import "strconv"

type Cache struct {
	Channel_ID int
	Data       map[string]*CachedItem
}

// Set a new item in the cache
func (c *Cache) Set(key string, value interface{}, ttl int) {
	LOGGER.Info("Setting key: " + key)
	MUTEX.Lock()
	defer MUTEX.Unlock()
	item := CachedItem{Key: key, Value: value, TTL: ttl}
	c.Data[key] = &item
}

// Get an item from the cache
func (c *Cache) Get(key string) interface{} {
	LOGGER.Info("Getting key: " + key)
	MUTEX.Lock()
	defer MUTEX.Unlock()
	item, ok := c.Data[key]
	if !ok {
		return nil
	}
	return item.Value
}

// Delete an item from the cache
func (c *Cache) Delete(key string) bool {
	LOGGER.Warning("Deleting key: " + key)
	MUTEX.Lock()
	defer MUTEX.Unlock()
	item, ok := c.Data[key]
	if !ok {
		LOGGER.Warning("Key not found: " + key)
		return false
	}
	delete(c.Data, item.Key)
	return true
}

// Delete all items from the cache
func (c *Cache) DeleteAll() {
	LOGGER.Warning("Deleting all keys from cache: " + strconv.Itoa(c.Channel_ID))
	MUTEX.Lock()
	c.Data = nil
	MUTEX.Unlock()
}

// Get size of the cache
func (c *Cache) GetSize() int {
	LOGGER.Info("Getting size of cache: " + strconv.Itoa(c.Channel_ID))
	return len(c.Data)
}

// Get all keys from the cache
func (c *Cache) GetKeys() []string {
	LOGGER.Info("Getting keys from cache: " + strconv.Itoa(c.Channel_ID))
	var keys []string
	for key := range c.Data {
		keys = append(keys, key)
	}
	return keys
}

// Verify if a key exists in the cache
func (c *Cache) HasKey(key string) bool {
	LOGGER.Info("Checking if key exists in cache: " + strconv.Itoa(c.Channel_ID))
	MUTEX.Lock()
	defer MUTEX.Unlock()
	_, ok := c.Data[key]
	return ok
}

// Get item TTL from the cache
func (c *Cache) ItemTTL(key string) int {
	LOGGER.Info("Getting item TTL from cache: " + strconv.Itoa(c.Channel_ID))
	MUTEX.Lock()
	defer MUTEX.Unlock()

	item, ok := c.Data[key]
	if !ok {
		return -1
	}
	// If the item is expired, delete it and return -1
	if item.IsExpired() {
		delete(c.Data, item.Key)
		return -1
	}
	return item.TTL
}
