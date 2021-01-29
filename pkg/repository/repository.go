// Package repository provides wrapper for go-pg package to simplify unit testing
package repository

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Repository list of repository interfaces
type Repository interface {
	SelectOrCreator
	Creator
	Updater
	Finder
	Deleter
	Transactor
	Executor
}

// SelectOrCreator interface to select or create new model
type SelectOrCreator interface {
	SelectOrCreate(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) (bool, error)
}

// Creator interface for creating models
type Creator interface {
	Create(ctx context.Context, model interface{}, values ...interface{}) error
}

// Updater updater interface for updating the model int database
type Updater interface {
	Update(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, fields ...interface{}) error
}

// Finder find model inside the repo
type Finder interface {
	Find(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) error
}

// Deleter delete model from the repo
type Deleter interface {
	Delete(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) error
}

// Transactor run set of operations in transaction
type Transactor interface {
	Transaction(ctx context.Context, callback func(db *pg.Tx) error) error
}

// Executor execute db query
type Executor interface {
	Exec(ctx context.Context, query string, params ...interface{}) error
}
