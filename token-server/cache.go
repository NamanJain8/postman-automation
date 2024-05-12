package main

import (
	"fmt"
	"sync"
)

type cacheT struct {
	sync.RWMutex
	tokenCache map[string]string
}

var cache cacheT

func init() {
	cache = cacheT{tokenCache: make(map[string]string)}
}

func key(r TokenRequest) string {
	return fmt.Sprintf("%s", r.Username)
}

func (c *cacheT) get(r TokenRequest) string {
	c.RLock()
	defer c.RUnlock()
	return c.tokenCache[key(r)]
}

func (c *cacheT) set(r TokenRequest, token string) {
	c.Lock()
	defer c.Unlock()
	c.tokenCache[key(r)] = token
}
