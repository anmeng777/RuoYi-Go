// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go
// Email: hot_kun@hotmail.com or BusinessCallKun@gmail.com

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"RuoYi-Go/internal/domain/model"
)

func newSysPost(db *gorm.DB, opts ...gen.DOOption) sysPost {
	_sysPost := sysPost{}

	_sysPost.sysPostDo.UseDB(db, opts...)
	_sysPost.sysPostDo.UseModel(&model.SysPost{})

	tableName := _sysPost.sysPostDo.TableName()
	_sysPost.ALL = field.NewAsterisk(tableName)
	_sysPost.PostID = field.NewInt64(tableName, "post_id")
	_sysPost.PostCode = field.NewString(tableName, "post_code")
	_sysPost.PostName = field.NewString(tableName, "post_name")
	_sysPost.PostSort = field.NewInt32(tableName, "post_sort")
	_sysPost.Status = field.NewString(tableName, "status")
	_sysPost.CreateBy = field.NewString(tableName, "create_by")
	_sysPost.CreateTime = field.NewTime(tableName, "create_time")
	_sysPost.UpdateBy = field.NewString(tableName, "update_by")
	_sysPost.UpdateTime = field.NewTime(tableName, "update_time")
	_sysPost.Remark = field.NewString(tableName, "remark")

	_sysPost.fillFieldMap()

	return _sysPost
}

type sysPost struct {
	sysPostDo sysPostDo

	ALL        field.Asterisk
	PostID     field.Int64  // 岗位ID
	PostCode   field.String // 岗位编码
	PostName   field.String // 岗位名称
	PostSort   field.Int32  // 显示顺序
	Status     field.String // 状态（0正常 1停用）
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间
	Remark     field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysPost) Table(newTableName string) *sysPost {
	s.sysPostDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysPost) As(alias string) *sysPost {
	s.sysPostDo.DO = *(s.sysPostDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysPost) updateTableName(table string) *sysPost {
	s.ALL = field.NewAsterisk(table)
	s.PostID = field.NewInt64(table, "post_id")
	s.PostCode = field.NewString(table, "post_code")
	s.PostName = field.NewString(table, "post_name")
	s.PostSort = field.NewInt32(table, "post_sort")
	s.Status = field.NewString(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysPost) WithContext(ctx context.Context) *sysPostDo { return s.sysPostDo.WithContext(ctx) }

func (s sysPost) TableName() string { return s.sysPostDo.TableName() }

func (s sysPost) Alias() string { return s.sysPostDo.Alias() }

func (s sysPost) Columns(cols ...field.Expr) gen.Columns { return s.sysPostDo.Columns(cols...) }

func (s *sysPost) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysPost) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["post_id"] = s.PostID
	s.fieldMap["post_code"] = s.PostCode
	s.fieldMap["post_name"] = s.PostName
	s.fieldMap["post_sort"] = s.PostSort
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysPost) clone(db *gorm.DB) sysPost {
	s.sysPostDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysPost) replaceDB(db *gorm.DB) sysPost {
	s.sysPostDo.ReplaceDB(db)
	return s
}

type sysPostDo struct{ gen.DO }

func (s sysPostDo) Debug() *sysPostDo {
	return s.withDO(s.DO.Debug())
}

func (s sysPostDo) WithContext(ctx context.Context) *sysPostDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysPostDo) ReadDB() *sysPostDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysPostDo) WriteDB() *sysPostDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysPostDo) Session(config *gorm.Session) *sysPostDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysPostDo) Clauses(conds ...clause.Expression) *sysPostDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysPostDo) Returning(value interface{}, columns ...string) *sysPostDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysPostDo) Not(conds ...gen.Condition) *sysPostDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysPostDo) Or(conds ...gen.Condition) *sysPostDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysPostDo) Select(conds ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysPostDo) Where(conds ...gen.Condition) *sysPostDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysPostDo) Order(conds ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysPostDo) Distinct(cols ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysPostDo) Omit(cols ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysPostDo) Join(table schema.Tabler, on ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysPostDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysPostDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysPostDo) Group(cols ...field.Expr) *sysPostDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysPostDo) Having(conds ...gen.Condition) *sysPostDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysPostDo) Limit(limit int) *sysPostDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysPostDo) Offset(offset int) *sysPostDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysPostDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysPostDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysPostDo) Unscoped() *sysPostDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysPostDo) Create(values ...*model.SysPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysPostDo) CreateInBatches(values []*model.SysPost, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysPostDo) Save(values ...*model.SysPost) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysPostDo) First() (*model.SysPost, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPost), nil
	}
}

func (s sysPostDo) Take() (*model.SysPost, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPost), nil
	}
}

func (s sysPostDo) Last() (*model.SysPost, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPost), nil
	}
}

func (s sysPostDo) Find() ([]*model.SysPost, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysPost), err
}

func (s sysPostDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysPost, err error) {
	buf := make([]*model.SysPost, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysPostDo) FindInBatches(result *[]*model.SysPost, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysPostDo) Attrs(attrs ...field.AssignExpr) *sysPostDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysPostDo) Assign(attrs ...field.AssignExpr) *sysPostDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysPostDo) Joins(fields ...field.RelationField) *sysPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysPostDo) Preload(fields ...field.RelationField) *sysPostDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysPostDo) FirstOrInit() (*model.SysPost, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPost), nil
	}
}

func (s sysPostDo) FirstOrCreate() (*model.SysPost, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysPost), nil
	}
}

func (s sysPostDo) FindByPage(offset int, limit int) (result []*model.SysPost, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysPostDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysPostDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysPostDo) Delete(models ...*model.SysPost) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysPostDo) withDO(do gen.Dao) *sysPostDo {
	s.DO = *do.(*gen.DO)
	return s
}
