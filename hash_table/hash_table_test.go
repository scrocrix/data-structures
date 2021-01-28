package hash_table_test

import (
	"github.com/scrocrix/data-structures/hash_table"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type hashTableTest struct {
	suite.Suite
}

func (unit *hashTableTest) TestInsert() {
	unit.Run("Return True When Successful", func() {
		sut := hash_table.HashTable{
			Bucket: map[int][]int{},
		}

		require.True(unit.T(), sut.Insert(20))

		require.NotEmpty(unit.T(), sut.Bucket)
		require.Equal(unit.T(), 1, len(sut.Bucket))
		require.Equal(unit.T(), 20, sut.Bucket[6][0])
	})

	unit.Run("Append key element to a linked list when collision happens", func() {
		sut := hash_table.HashTable{
			Bucket: map[int][]int{},
		}

		require.True(unit.T(), sut.Insert(50))

		require.NotEmpty(unit.T(), sut.Bucket)
		require.Equal(unit.T(), 1, len(sut.Bucket))
		require.Equal(unit.T(), 1, len(sut.Bucket[1]))

		require.True(unit.T(), sut.Insert(85))
		require.Equal(unit.T(), 1, len(sut.Bucket))
		require.Equal(unit.T(), 2, len(sut.Bucket[1]))
	})
}

func (unit *hashTableTest) TestSearch() {
	unit.Run("Return the element from within the Hash Table", func() {
		sut := hash_table.HashTable{
			Bucket: map[int][]int{},
		}

		require.True(unit.T(), sut.Insert(50))
		require.True(unit.T(), sut.Insert(85))

		result, err := sut.Search(50)

		require.Nil(unit.T(), err)

		require.Equal(unit.T(), 50, result)
	})

	unit.Run("Return negative one when element was not found", func() {
		sut := hash_table.HashTable{
			Bucket: map[int][]int{},
		}

		require.True(unit.T(), sut.Insert(500))

		result, err := sut.Search(1)

		require.NotNil(unit.T(), err)

		require.Equal(unit.T(), hash_table.ErrNoSuchKeyFound.Error(), err.Error())

		require.Equal(unit.T(), -1, result)
	})
}

func TestHashTable(t *testing.T) {
	suite.Run(t, new(hashTableTest))
}
