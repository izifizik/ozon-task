package cache

type cache struct {
	cache map[string]string
}

//NewCache - create new Cache interface reference
func NewCache() Cache {
	return &cache{cache: make(map[string]string)}
}

//Set - set value to cache (map)
func (c *cache) Set(uuid string, value string) {
	c.cache[uuid] = value
}

//Get - get value from cache (map)
func (c *cache) Get(uuid string) (string, bool) {
	value, ok := c.cache[uuid]
	return value, ok
}
