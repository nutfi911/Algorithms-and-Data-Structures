package hashmap

import (
	"math"
)

var privateKey, prime int = 96, 31

type Data struct {
	Key   string
	Value string
}
type HashTable struct {
	Size   int
	KeyMap [][]Data
}

func New(size int) *HashTable {
	return &HashTable{Size: size, KeyMap: make([][]Data, size)}
}

func hash(key string, arraySize int) int {
	var total int = 0

	for i := 0; i < int(math.Min(float64(len(key)), 100.)); i++ {
		char := string(key[i])
		var v int = int(char[0]) - privateKey
		total = (total + prime + v) % arraySize
	}
	return int(math.Abs(float64(total)))
}

func (ht *HashTable) Set(key string, value string) int {
	index := hash(key, ht.Size)

	if ht.KeyMap[index] == nil {
		ht.KeyMap[index] = []Data{}
	}
	ht.KeyMap[index] = append(ht.KeyMap[index], Data{Key: key, Value: value})
	return index
}

func (ht *HashTable) Get(key string) (value string) {
	index := hash(key, ht.Size)

	if ht.KeyMap[index] != nil {
		for i := 0; i < len(ht.KeyMap[index]); i++ {
			if ht.KeyMap[index][i].Key == key {
				return ht.KeyMap[index][i].Value
			}
		}
	}
	return ""
}

func (ht *HashTable) Keys() []string {
	keys := []string{}

	for i := 0; i < len(ht.KeyMap); i++ {
		for j := 0; j < len(ht.KeyMap[i]); j++ {
			keys = append(keys, ht.KeyMap[i][j].Key)
		}
	}

	return keys
}

func (ht *HashTable) Values() []string {
	values := []string{}

	for i := 0; i < len(ht.KeyMap); i++ {
		for j := 0; j < len(ht.KeyMap[i]); j++ {
			values = append(values, ht.KeyMap[i][j].Value)
		}
	}

	return values
}
