package maps

import "errors"
type Dictionary map[string]string

func Search(m map[string]string, key string) string{
	return  m[key]
}


var ErrNotFound = errors.New("could not find the word you were looking for")
func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string){
	d[key] = value
}