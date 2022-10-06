package main

import (
	"encoding/gob"
	"errors"
	"os"
	"time"
)

type CacheList struct {
	Caches map[int]*Cache
}

type SavedItem struct {
	Key        string
	Data       interface{}
	TIME_SAVED int
	TTL        int
}

type SaveList struct {
	Saves map[int]*SavedItem
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

func (c *CacheList) SaveCaches() {
	data := SaveList{Saves: make(map[int]*SavedItem)}
	for _, cache := range c.Caches {
		for key, item := range cache.Data {
			nowtime := int(time.Now().Unix())
			data.Saves[cache.Channel_ID] = &SavedItem{Key: key, Data: item.Value, TIME_SAVED: nowtime, TTL: item.TTL}
		}
	}
	data.Save("cache.extra")
}

func (c *CacheList) LoadCaches() {
	data := SaveList{Saves: make(map[int]*SavedItem)}
	data.Load("cache.extra")
	for key, item := range data.Saves {
		nowtime := int(time.Now().Unix())
		ttl := item.TTL - (nowtime - item.TIME_SAVED)
		if ttl > 0 {
			c.SetItemInCache(key, item.Key, item.Data, ttl)
		}
	}
	os.Remove("cache.extra")
}

func (sl *SaveList) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		err = errors.New("An error has occurred trying to save the cache to disk." + err.Error())
		LOGGER.Error(err.Error())
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(sl)
	if err != nil {
		err = errors.New("An error has occurred trying to save the cache to disk." + err.Error())
		LOGGER.Error(err.Error())
		return err
	}
	return nil
}

func (sl *SaveList) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		err = errors.New("An error has occurred trying to load the cache from disk." + err.Error())
		LOGGER.Error(err.Error())
		return err
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(sl)
	if err != nil {
		err = errors.New("An error has occurred trying to load the cache from disk." + err.Error())
		LOGGER.Error(err.Error())
		return err
	}
	return nil
}
