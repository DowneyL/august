package migration

import (
	"fmt"
	"github.com/DowneyL/august/model/migration/tmpl"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func repoMigrate(option Option) error {
	if option.Ro.NotRun {
		return nil
	}

	if err := os.MkdirAll(option.Ro.To, os.ModePerm); err != nil {
		return err
	}

	if err := createDBIfNotExists(option); err != nil {
		return err
	}

	if err := createBaseRepoIfNotExists(option); err != nil {
		return err
	}

	files, err := readPoFiles(option)
	if err != nil {
		return err
	}

	for _, f := range files {
		option.Ro.StructName = camelcase(basename(f.Name()))
		err = runRepoMigrate(option, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateFile(fp, tpl string, option Option) error {
	if _, err := os.Stat(fp); err == nil && !option.Ro.Overwrite {
		return err
	}

	fd, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer func() {
		_ = fd.Close()
	}()

	t, err2 := template.New("").Parse(tpl)
	if err2 != nil {
		return err2
	}

	return t.Execute(fd, option)
}

func createDBIfNotExists(option Option) error {
	dbFilePath := fmt.Sprintf("%s/%s", option.Do.To, "db.go")
	return generateFile(dbFilePath, tmpl.DBTemplate, option)
}

func createBaseRepoIfNotExists(option Option) error {
	baseFilepath := fmt.Sprintf("%s/%s", option.Ro.To, "base.go")
	return generateFile(baseFilepath, tmpl.BaseTemplate, option)
}

func runRepoMigrate(option Option, f fs.FileInfo) error {
	fp := fmt.Sprintf("%s/%s", option.Ro.To, f.Name())
	return generateFile(fp, tmpl.RepoTemplate, option)
}

func contains(elems []string, elem string) bool {
	for _, e := range elems {
		if elem == e {
			return true
		}
	}
	return false
}

func basename(fp string) string {
	fb := filepath.Base(fp)
	ext := filepath.Ext(fb)
	return strings.TrimSuffix(fb, ext)
}

func camelcase(snakeCase string) string {
	split := strings.Split(snakeCase, "_")
	for k, v := range split {
		b := []byte(v)
		if b[0] >= 'a' && b[0] <= 'z' {
			b[0] = b[0] - ('a' - 'A')
		}
		split[k] = string(b)
	}

	return strings.Join(split, "")
}

func readPoFiles(option Option) ([]fs.FileInfo, error) {
	res, err := ioutil.ReadDir(option.Ro.From)
	if err != nil {
		return nil, err
	}

	if !option.IsInit() {
		var fres []fs.FileInfo
		tables := option.Tables()
		for _, v := range res {
			if contains(tables, basename(v.Name())) {
				fres = append(fres, v)
			}
		}
		return fres, nil
	}

	return res, err
}
