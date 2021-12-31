package tmpl

var BaseTemplate = `package repo

import (
	"{{.ModuleName}}/{{.Do.To}}"
	"github.com/DowneyL/august/model/repo"
)

type {{.Ro.RepositoryName}} struct {
	repo.BaseRepository
}

func {{.Ro.NewMethodName}}(tableName string) repository {
	return repository{
		repo.NewRepository(db.Connect(), tableName),
	}
}
`
