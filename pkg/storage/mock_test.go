package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStorage(store Storage) error {
	return nil
}

func TestMock(t *testing.T) {
	assert := assert.New(t)
	mock := NewMock()
	assert.NotNil(mock)
	assert.Nil(testStorage(mock))
}
