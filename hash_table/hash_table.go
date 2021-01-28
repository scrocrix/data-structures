package hash_table

import (
	"errors"
	"sort"
)

var ErrNoSuchKeyFound = errors.New("error: no such key was found")

// HashTable is a data structure which is divided into three parts: Key, Hash Function and the Bucket.
//
// The advantage of using Hash Tables is the ability to deal with large arbitrary data sets in O(1) time.
type HashTable struct {
	// This bucket is mapping an Map Index to a "slice of integers" (a.k.a. Linked List Structure)
	Bucket map[int][]int
}

// hash contains the algorithm responsible for converting our key value into a hash value.
func (ht *HashTable) hash(key int) int {
	return key % 7
}

// Insert is responsible for storing the input key value into the Hash Table's data bucket.
func (ht *HashTable) Insert(key int) bool {
	hv := ht.hash(key)

	ht.Bucket[hv] = append(ht.Bucket[hv], key)

	sort.Ints(ht.Bucket[hv])

	return true
}

// Search is responsible for finding the element in the Hash Table.
func (ht *HashTable) Search(key int) (int, error) {
	ll := ht.Bucket[ht.hash(key)]

	l := 0
	r := len(ll) - 1

	for l <= r {
		target := l + r/2

		if ll[target] == key {
			return ll[target], nil
		}

		if ll[target] < key {
			l = target + 1
		}

		if ll[target] > key {
			r = target - 1
		}
	}

	return -1, ErrNoSuchKeyFound
}
