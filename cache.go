package cache

var Cache []map[string]interface{}

func New() *[]map[string]interface{} {
	Cache := make([]map[string]interface{}, 0)
	return &Cache
}

func Set(key string, value interface{}) {
	Cache = append(Cache, map[string]interface{}{key: value})
}
func Get(key string) interface{} {
	for _, i := range Cache {
		for k, v := range i {
			if k == key {
				return v
			}
		}
	}
	return nil
}
