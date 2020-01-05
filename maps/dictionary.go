package maps

import "errors"
type Dictionary map[string]string

func Search(m map[string]string, key string) string{
	return  m[key]
}


const (
    ErrNotFound   = DictionaryErr("could not find the word you were looking for")
    ErrWordExists = DictionaryErr("cannot add word because it already exists")
)


type DictionaryErr string

func (e DictionaryErr) Error() string{
	return string(e)
}



func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)
	switch err{
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	
	return nil
}