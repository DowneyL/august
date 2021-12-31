package repo

import (
	"context"
	"gorm.io/gorm"
)

type Repository interface {
	WithContext(ctx context.Context) *gorm.DB
	Table(ctx context.Context) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Or(query interface{}, args ...interface{}) *gorm.DB
}

type BaseRepository struct {
	db        *gorm.DB
	TableName string
}

func NewRepository(db *gorm.DB, tableName string) BaseRepository {
	return BaseRepository{
		db:        db,
		TableName: tableName,
	}
}

func (bs BaseRepository) WithContext(ctx context.Context) *gorm.DB {
	return bs.db.WithContext(ctx)
}

func (bs BaseRepository) Table(ctx context.Context) *gorm.DB {
	return bs.WithContext(ctx).Table(bs.TableName)
}

func (bs BaseRepository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return bs.db.Table(bs.TableName).Where(query, args...)
}

func (bs BaseRepository) Or(query interface{}, args ...interface{}) *gorm.DB {
	return bs.db.Table(bs.TableName).Or(query, args...)
}
