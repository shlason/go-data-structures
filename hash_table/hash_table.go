package hashTable

type element interface{}

type hashTable struct {
	table []element
}

func djb2HashCode(key string) int {
	hash := 5381
	for _, val := range []byte(key) {
		hash = hash*33 + int(val)
	}
	return hash % 1013
}

func New() *hashTable {
	t := new(hashTable)
	t.table = make([]element, 8)
	return t
}

func (ht *hashTable) Put(key string, e element) {
	k := djb2HashCode(key)
	ht.table[k] = e
}

func (ht *hashTable) Get(key string) element {
	return ht.table[djb2HashCode(key)]
}

func (ht *hashTable) Remove(key string) {
	ht.table[djb2HashCode(key)] = nil
}
