package db

import (
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/protsack-stephan/dev-toolkit/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func testRepository(repo repository.Repository) error {
	return nil
}

func TestRepository(t *testing.T) {
	assert := assert.New(t)

	t.Run("repository interface match", func(t *testing.T) {
		assert.Nil(testRepository(NewRepository(new(pg.DB))))
	})
}
