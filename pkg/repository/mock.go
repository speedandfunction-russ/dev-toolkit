package repository

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// NewMock create new repository for testing
func NewMock() *Mock {
	return new(Mock)
}

// Mock is short for repository
type Mock struct{}

// SelectOrCreate select new language and create if not exists
func (Mock) SelectOrCreate(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) (bool, error) {
	return true, nil
}

// Create make new model inside the database
func (Mock) Create(ctx context.Context, model interface{}, values ...interface{}) (orm.Result, error) {
	return nil, nil
}

// Update update fields for the model
func (Mock) Update(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, fields ...interface{}) (orm.Result, error) {
	return nil, nil
}

// Find find the model in database
func (Mock) Find(ctx context.Context, model interface{}, modifier func(q *orm.Query) *orm.Query, values ...interface{}) error {
	return nil
}

// Delete delete the model from database
func (Mock) Delete(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) (orm.Result, error) {
	return nil, nil
}

// Transaction run set of queries in transaction
func (Mock) Transaction(ctx context.Context, callback func(db *pg.Tx) error) error {
	return nil
}

// Exec run query on the repository
func (Mock) Exec(ctx context.Context, query string, params ...interface{}) (orm.Result, error) {
	return nil, nil
}
