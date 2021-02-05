package s3

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/protsack-stephan/dev-toolkit/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func testStorage(store storage.Storage) error {
	return nil
}

func TestStorage(t *testing.T) {
	assert := assert.New(t)
	ses := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("1234", "5678", ""),
	}))

	store := NewStorage(ses, "new")
	assert.NotNil(store)
	assert.Nil(testStorage(store))
}
