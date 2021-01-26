package hash_table

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

	if len(ht.Bucket[hv]) > 0 {
		ht.Bucket[hv] = append(ht.Bucket[hv], key)
	} else {
		ht.Bucket[hv] = append(ht.Bucket[hv], key)
	}

	return true
}
