package types

type HashMap map[interface{}]interface{}

func (h HashMap) Get(key interface{}) interface{} {
	return h[key]
}

func (h HashMap) Exists(key interface{}) bool {
	_, exists := h[key]
	return exists
}

func (h HashMap) Set(key interface{}, value interface{}) {
	h[key] = value
}
