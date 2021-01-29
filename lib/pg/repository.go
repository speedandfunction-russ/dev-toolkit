package pg

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Repository is short for repository
type Repository struct {
	db *pg.DB
}

// NewRepository create new repository instance
func NewRepository(conn *pg.DB) *Repository {
	return &Repository{
		conn,
	}
}

// Create make new model inside the database
func (r *Repository) Create(ctx context.Context, model interface{}, values ...interface{}) error {
	_, err := r.db.WithContext(ctx).Model(model).Insert(values...)
	return err
}

// SelectOrCreate select model from db and create if not exists
func (r *Repository) SelectOrCreate(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) (bool, error) {
	return modifier(r.db.WithContext(ctx).Model(model)).SelectOrInsert(values...)
}

// Update update fields for the model
func (r *Repository) Update(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, fields ...interface{}) error {
	_, err := modifier(r.db.WithContext(ctx).Model(model)).Update(fields...)
	return err
}

// Find find the model in database
func (r *Repository) Find(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) error {
	return modifier(r.db.WithContext(ctx).Model(model)).Select(values...)
}

// Delete delete model from database
func (r *Repository) Delete(ctx context.Context, model interface{}, modifier func(*orm.Query) *orm.Query, values ...interface{}) error {
	_, err := modifier(r.db.WithContext(ctx).Model(model)).Delete(values...)
	return err
}

// Exec run query on the database
func (r *Repository) Exec(ctx context.Context, query string, params ...interface{}) error {
	_, err := r.db.WithContext(ctx).Exec(query, params...)
	return err
}

// Transaction run set of queries in transaction
func (r *Repository) Transaction(ctx context.Context, callback func(db *pg.Tx) error) error {
	return r.db.RunInTransaction(ctx, callback)
}
