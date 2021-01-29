package fs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const storageTestVol = "./testdata"
const storageTestPath = "test.txt"
const storageTestWalkPath = "testwalk.txt"

var storageTestExpire = time.Second * 1
var storageTestData = []byte("hello storage")

func TestStorage(t *testing.T) {
	store := NewStorage(storageTestVol)
	assert := assert.New(t)

	t.Run("walk path", func(t *testing.T) {
		assert.NoError(store.Walk("/", func(path string) {
			assert.Equal(storageTestWalkPath, path)
		}))
	})

	t.Run("create file", func(t *testing.T) {
		file, err := store.Create(storageTestPath)
		assert.NoError(err)
		file.Close()
		assert.NoError(os.Remove(fmt.Sprintf("%s/%s", storageTestVol, storageTestPath)))
	})

	t.Run("put file", func(t *testing.T) {
		assert.NoError(store.Put(storageTestPath, bytes.NewReader(storageTestData)))
	})

	t.Run("get file", func(t *testing.T) {
		body, err := store.Get(storageTestPath)
		assert.NoError(err)
		defer body.Close()

		data, err := ioutil.ReadAll(body)
		assert.NoError(err)
		assert.Equal(storageTestData, data)
	})

	t.Run("link file", func(t *testing.T) {
		loc, err := store.Link(storageTestPath, storageTestExpire)
		assert.NoError(err)
		assert.Equal(fmt.Sprintf("%s/%s", storageTestVol, storageTestPath), loc)
	})

	t.Run("delete file", func(t *testing.T) {
		assert.NoError(store.Delete(storageTestPath))
	})
}
