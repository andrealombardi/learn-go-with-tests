package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExist = DictionaryErr("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrAlreadyExist
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	d[key] = value
	return nil
}
