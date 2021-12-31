package tmpl

var RepoTemplate = `package repo

import (
	"context"
	"{{.ModuleName}}/{{.Ro.From}}"
	"github.com/DowneyL/august/model/repo"
)

type {{.Ro.StructName}} struct {
	Instance {{.Ro.RepositoryName}}
}

func New{{.Ro.StructName}}() {{.Ro.StructName}} {
	var pov *po.{{.Ro.StructName}}
	return {{.Ro.StructName}}{ {{.Ro.NewMethodName}}(pov.TableName()) }
}

func (r {{.Ro.StructName}}) All(ctx context.Context) ([]*po.{{.Ro.StructName}}, error) {
	return r.FindWhere(ctx)
}

func (r {{.Ro.StructName}}) Count(ctx context.Context, conds ...interface{}) (count int64, err error) {
	tx := r.Instance.Table(ctx)
	if len(conds) > 0 {
		tx.Where(conds[0], conds[1:]...)
	}
	err = tx.Count(&count).Error
	return
}

func (r {{.Ro.StructName}}) First(ctx context.Context, conds ...interface{}) (res *po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).First(&res, conds...).Error
	return
}

func (r {{.Ro.StructName}}) Last(ctx context.Context, conds ...interface{}) (res *po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Last(&res, conds...).Error
	return
}

func (r {{.Ro.StructName}}) Find(ctx context.Context, id int64) (res *po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Find(&res, id).Error
	return
}

func (r {{.Ro.StructName}}) FindByField(
	ctx context.Context, query string, arg interface{}) (res []*po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Find(&res, query, arg).Error
	return
}

func (r {{.Ro.StructName}}) FindWhere(
	ctx context.Context, conds ...interface{}) (res []*po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Find(&res, conds...).Error
	return
}

func (r {{.Ro.StructName}}) FindWhereIn(
	ctx context.Context, query string, arg []interface{}) (res []*po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Find(&res, query, arg).Error
	return
}

func (r {{.Ro.StructName}}) FindWhereNotIn(
	ctx context.Context, query string, arg []interface{}) (res []*po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Not(query, arg).Find(&res).Error
	return
}

func (r {{.Ro.StructName}}) FindWhereBetween(
	ctx context.Context, query string, arg []interface{}) (res []*po.{{.Ro.StructName}}, err error) {
	if err = repo.IsBetweenQuery(&query, arg); err != nil {
		return
	}
	err = r.Instance.Table(ctx).Find(&res, query, arg[0], arg[1]).Error
	return
}

func (r {{.Ro.StructName}}) Limit(ctx context.Context, limit int) (res []*po.{{.Ro.StructName}}, err error) {
	err = r.Instance.Table(ctx).Limit(limit).Find(&res).Error
	return
}

func (r {{.Ro.StructName}}) Paginate(ctx context.Context, currentPage, limit int) (res []*po.{{.Ro.StructName}}, err error) {
	paginator := repo.NewPaginator(currentPage, limit)
	err = r.Instance.Table(ctx).Scopes(paginator.Paginate()).Find(&res).Error
	return
}

func (r {{.Ro.StructName}}) Create(ctx context.Context, data ...*po.{{.Ro.StructName}}) (int64, error) {
	res := r.Instance.Table(ctx).Create(data)
	return res.RowsAffected, res.Error
}

func (r {{.Ro.StructName}}) UpdateWhere(
	ctx context.Context, updates map[string]interface{}, query interface{}, args ...interface{}) (int64, error) {
	res := r.Instance.Table(ctx).Where(query, args).Updates(updates)

	return res.RowsAffected, res.Error
}

func (r {{.Ro.StructName}}) Update(ctx context.Context, id int64, updates map[string]interface{}) (int64, error) {
	return r.UpdateWhere(ctx, updates, r.Instance.Where("id", id))
}

func (r {{.Ro.StructName}}) DeleteWhere(
	ctx context.Context, query interface{}, args ...interface{}) (int64, error) {
	args = append(args, query)
	db := r.Instance.Table(ctx).Delete(nil, args...)
	return db.RowsAffected, db.Error
}

func (r {{.Ro.StructName}}) Delete(ctx context.Context, id int64) (int64, error) {
	return r.DeleteWhere(ctx, r.Instance.Where("id", id))
}`
