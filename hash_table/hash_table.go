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

// search performs a binary search in the linked list.
//
// It receives an argument which is a high-order function to manipulate custom scenarios.
func (ht *HashTable) search(key int, ll []int, isFound func(target int) int) int {
	l := 0
	r := len(ll) - 1

	for l <= r {
		target := l + r/2

		if ll[target] == key {
			return isFound(target)
		}

		if ll[target] < key {
			l = target + 1
		}

		if ll[target] > key {
			r = target - 1
		}
	}

	return -1
}

// Delete is responsible for searching the key and deleting it once it exists.
func (ht *HashTable) Delete(key int) error {
	hv := ht.hash(key)
	ll := ht.Bucket[hv]

	ht.search(key, ll, func(target int) int {
		// we delete the item from the linked list
		ht.Bucket[hv] = append(ht.Bucket[hv][:target], ht.Bucket[hv][target+1:]...)

		// we delete the index reference from the hash table, if the linked list has no items left
		if len(ht.Bucket[ht.hash(key)]) == 0 {
			delete(ht.Bucket, hv)
		}

		return key
	})

	return nil
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

	n := ht.search(key, ll, func(target int) int {
		return key
	})

	if n == -1 {
		return n, ErrNoSuchKeyFound
	}

	return n, nil
}
