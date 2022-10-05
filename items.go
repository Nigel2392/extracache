package main

type CachedItem struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	TTL   int         `json:"ttl"`
}

// Verify if item is expired.
func (c *CachedItem) IsExpired() bool {
	return c.TTL <= 0
}
