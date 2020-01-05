package maps

type Dictionary map[string]string

func Search(m map[string]string, key string) string{
	return  m[key]
}

func (d Dictionary) Search(key string) string {
	return d[key]
}