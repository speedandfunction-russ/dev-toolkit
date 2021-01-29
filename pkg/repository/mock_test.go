package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRepository(repo Repository) error {
	return nil
}

func TestMock(t *testing.T) {
	mock := NewMock()
	assert := assert.New(t)

	assert.NotNil(mock)
	assert.Nil(testRepository(mock))
}
