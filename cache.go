package cache

type Cache []map[string]interface{}

var cache Cache

func New() Cache {
	cache := make([]map[string]interface{}, 0)
	return cache
}

func (c Cache) Set(key string, value interface{}) {
	cache = append(cache, map[string]interface{}{key: value})
}
func Get(key string) interface{} {
	for _, i := range cache {
		for k, v := range i {
			if k == key {
				return v
			}
		}
	}
	return nil
}
