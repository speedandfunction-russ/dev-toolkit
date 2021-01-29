// Package s3 this package in intention to hide aws s3 storage implementation
// under the interface that will give you the ability to user other cloud providers
// in the future
package s3

import (
	"errors"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const maxUploadParts = 20000
const partSize = 1024 * 1024 * 5 * 2

// NewStorage create new storage instance
func NewStorage(ses *session.Session, bucket string) *Storage {
	return &Storage{
		s3:     s3.New(ses),
		bucket: bucket,
		uploader: s3manager.NewUploader(ses, func(upl *s3manager.Uploader) {
			upl.MaxUploadParts = maxUploadParts
			upl.PartSize = partSize
		}),
	}
}

// Storage interface adaptation for s3
type Storage struct {
	bucket   string
	uploader *s3manager.Uploader
	s3       *s3.S3
}

// Walk recursively look for files in directory
func (s *Storage) Walk(path string, callback func(path string)) error {
	res, err := s.s3.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(path),
	})

	if err != nil {
		return err
	}

	for _, object := range res.Contents {
		callback(*object.Key)
	}

	return nil
}

// Create for create interface
func (s *Storage) Create(path string) (io.ReadWriteCloser, error) {
	return nil, errors.New("method unimplemented")
}

// Get file from s3 bucket
func (s *Storage) Get(path string) (io.ReadCloser, error) {
	out, err := s.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	return out.Body, err
}

// Put file into s3 bucket
func (s *Storage) Put(path string, body io.Reader) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   body,
	})

	return err
}

// Link generate expiration link for s3 access
func (s *Storage) Link(path string, expire time.Duration) (string, error) {
	req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	return req.Presign(expire)
}

// Delete remove object from s3
func (s *Storage) Delete(path string) error {
	_, err := s.s3.DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(path),
		Bucket: aws.String(s.bucket),
	})

	return err
}