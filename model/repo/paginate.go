package repo

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Paginator struct {
	Total       int64
	PageSize    int
	CurrentPage int
	Items       []interface{}
}

func NewPaginator(currentPage int, opts ...interface{}) Paginator {
	ps := viper.GetInt("page_size")
	if len(opts) > 0 {
		if nps, ok := opts[0].(int); ok {
			ps = nps
		}
	}

	return Paginator{
		PageSize:    ps,
		CurrentPage: currentPage,
	}
}

func (p Paginator) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset()).Limit(p.Limit())
	}
}

func (p Paginator) Offset() int {
	return (p.CurrentPage - 1) * p.PageSize
}

func (p Paginator) Limit() int {
	return p.PageSize
}
