package migration

import (
	"path"
	"path/filepath"
	"strings"
)

var (
	gormtPath = "pkg/model/migration/third_party/gormt"

	doToPathFunc PathFunc = func(app string) string {
		return "internal/" + app + "/model/db"
	}

	poToPathFunc PathFunc = func(app string) string {
		return "internal/" + app + "/model/po"
	}
)

type PathFunc func(app string) string

type RepoOption struct {
	From           string
	To             string
	StructName     string
	RepositoryName string
	NewMethodName  string
	Overwrite      bool
	NotRun         bool
}

type PoOption struct {
	GormtPath string
	To        string
	Overwrite bool
	NotRun    bool
}

type DBOption struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	To       string
}

func (opt PoOption) Gormt() string {
	if opt.GormtPath == "" {
		return "./gormt"
	}

	binPath := filepath.Clean(opt.GormtPath)
	return binPath + "/gormt"
}

type Option struct {
	AppName     string
	ModuleName  string
	TablePrefix string
	TableNames  string
	Ro          RepoOption
	Po          PoOption
	Do          DBOption
}

func (o Option) IsInit() bool {
	return o.TableNames == ""
}

func (o Option) Tables() []string {
	var res []string
	tablePrefix := o.TablePrefix
	hasMinusPrefix := strings.HasPrefix(tablePrefix, "-")
	for _, s := range strings.Split(o.TableNames, ",") {
		if hasMinusPrefix {
			res = append(res, strings.TrimPrefix(s, strings.TrimPrefix(tablePrefix, "-")))
		} else {
			res = append(res, tablePrefix+s)
		}
	}

	return res
}

func SetGormtPath(s string) {
	gormtPath = s
}

func SetPoToPathFunc(f PathFunc) {
	poToPathFunc = f
}

func SetDoToPathFunc(f PathFunc) {
	doToPathFunc = f
}

func NewDefaultOption(db DBOption, moduleName, app, tablePrefix, tableNames string) Option {
	po := PoOption{
		GormtPath: gormtPath,
		To:        poToPathFunc(app),
		NotRun:    false,
		Overwrite: true,
	}

	ro := RepoOption{
		From:           po.To,
		To:             path.Dir(po.To) + "/repo",
		RepositoryName: "repository",
		NewMethodName:  "newRepository",
		NotRun:         false,
		Overwrite:      true,
	}

	option := Option{
		AppName:     app,
		ModuleName:  moduleName,
		TablePrefix: tablePrefix,
		TableNames:  tableNames,
		Po:          po,
		Ro:          ro,
		Do:          db,
	}

	option.Do.To = doToPathFunc(app)

	return option
}
