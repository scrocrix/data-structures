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

func (unit *hashTableTest) TestInsertReturnTrueWhenSuccessful() {
	sut := hash_table.HashTable{
		Bucket: map[int][]int{},
	}

	require.True(unit.T(), sut.Insert(20))

	require.NotEmpty(unit.T(), sut.Bucket)
	require.Equal(unit.T(), 1, len(sut.Bucket))
	require.Equal(unit.T(), 20, sut.Bucket[6][0])
}

func TestHashTable(t *testing.T) {
	suite.Run(t, new(hashTableTest))
}
