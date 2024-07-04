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

func newSysLogininfor(db *gorm.DB, opts ...gen.DOOption) sysLogininfor {
	_sysLogininfor := sysLogininfor{}

	_sysLogininfor.sysLogininforDo.UseDB(db, opts...)
	_sysLogininfor.sysLogininforDo.UseModel(&model.SysLogininfor{})

	tableName := _sysLogininfor.sysLogininforDo.TableName()
	_sysLogininfor.ALL = field.NewAsterisk(tableName)
	_sysLogininfor.InfoID = field.NewInt64(tableName, "info_id")
	_sysLogininfor.UserName = field.NewString(tableName, "user_name")
	_sysLogininfor.Ipaddr = field.NewString(tableName, "ipaddr")
	_sysLogininfor.LoginLocation = field.NewString(tableName, "login_location")
	_sysLogininfor.Browser = field.NewString(tableName, "browser")
	_sysLogininfor.Os = field.NewString(tableName, "os")
	_sysLogininfor.Status = field.NewString(tableName, "status")
	_sysLogininfor.Msg = field.NewString(tableName, "msg")
	_sysLogininfor.LoginTime = field.NewTime(tableName, "login_time")

	_sysLogininfor.fillFieldMap()

	return _sysLogininfor
}

type sysLogininfor struct {
	sysLogininforDo sysLogininforDo

	ALL           field.Asterisk
	InfoID        field.Int64  // 访问ID
	UserName      field.String // 用户账号
	Ipaddr        field.String // 登录IP地址
	LoginLocation field.String // 登录地点
	Browser       field.String // 浏览器类型
	Os            field.String // 操作系统
	Status        field.String // 登录状态（0成功 1失败）
	Msg           field.String // 提示消息
	LoginTime     field.Time   // 访问时间

	fieldMap map[string]field.Expr
}

func (s sysLogininfor) Table(newTableName string) *sysLogininfor {
	s.sysLogininforDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysLogininfor) As(alias string) *sysLogininfor {
	s.sysLogininforDo.DO = *(s.sysLogininforDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysLogininfor) updateTableName(table string) *sysLogininfor {
	s.ALL = field.NewAsterisk(table)
	s.InfoID = field.NewInt64(table, "info_id")
	s.UserName = field.NewString(table, "user_name")
	s.Ipaddr = field.NewString(table, "ipaddr")
	s.LoginLocation = field.NewString(table, "login_location")
	s.Browser = field.NewString(table, "browser")
	s.Os = field.NewString(table, "os")
	s.Status = field.NewString(table, "status")
	s.Msg = field.NewString(table, "msg")
	s.LoginTime = field.NewTime(table, "login_time")

	s.fillFieldMap()

	return s
}

func (s *sysLogininfor) WithContext(ctx context.Context) *sysLogininforDo {
	return s.sysLogininforDo.WithContext(ctx)
}

func (s sysLogininfor) TableName() string { return s.sysLogininforDo.TableName() }

func (s sysLogininfor) Alias() string { return s.sysLogininforDo.Alias() }

func (s sysLogininfor) Columns(cols ...field.Expr) gen.Columns {
	return s.sysLogininforDo.Columns(cols...)
}

func (s *sysLogininfor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysLogininfor) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["info_id"] = s.InfoID
	s.fieldMap["user_name"] = s.UserName
	s.fieldMap["ipaddr"] = s.Ipaddr
	s.fieldMap["login_location"] = s.LoginLocation
	s.fieldMap["browser"] = s.Browser
	s.fieldMap["os"] = s.Os
	s.fieldMap["status"] = s.Status
	s.fieldMap["msg"] = s.Msg
	s.fieldMap["login_time"] = s.LoginTime
}

func (s sysLogininfor) clone(db *gorm.DB) sysLogininfor {
	s.sysLogininforDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysLogininfor) replaceDB(db *gorm.DB) sysLogininfor {
	s.sysLogininforDo.ReplaceDB(db)
	return s
}

type sysLogininforDo struct{ gen.DO }

func (s sysLogininforDo) Debug() *sysLogininforDo {
	return s.withDO(s.DO.Debug())
}

func (s sysLogininforDo) WithContext(ctx context.Context) *sysLogininforDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysLogininforDo) ReadDB() *sysLogininforDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysLogininforDo) WriteDB() *sysLogininforDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysLogininforDo) Session(config *gorm.Session) *sysLogininforDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysLogininforDo) Clauses(conds ...clause.Expression) *sysLogininforDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysLogininforDo) Returning(value interface{}, columns ...string) *sysLogininforDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysLogininforDo) Not(conds ...gen.Condition) *sysLogininforDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysLogininforDo) Or(conds ...gen.Condition) *sysLogininforDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysLogininforDo) Select(conds ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysLogininforDo) Where(conds ...gen.Condition) *sysLogininforDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysLogininforDo) Order(conds ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysLogininforDo) Distinct(cols ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysLogininforDo) Omit(cols ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysLogininforDo) Join(table schema.Tabler, on ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysLogininforDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysLogininforDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysLogininforDo) Group(cols ...field.Expr) *sysLogininforDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysLogininforDo) Having(conds ...gen.Condition) *sysLogininforDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysLogininforDo) Limit(limit int) *sysLogininforDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysLogininforDo) Offset(offset int) *sysLogininforDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysLogininforDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysLogininforDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysLogininforDo) Unscoped() *sysLogininforDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysLogininforDo) Create(values ...*model.SysLogininfor) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysLogininforDo) CreateInBatches(values []*model.SysLogininfor, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysLogininforDo) Save(values ...*model.SysLogininfor) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysLogininforDo) First() (*model.SysLogininfor, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLogininfor), nil
	}
}

func (s sysLogininforDo) Take() (*model.SysLogininfor, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLogininfor), nil
	}
}

func (s sysLogininforDo) Last() (*model.SysLogininfor, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLogininfor), nil
	}
}

func (s sysLogininforDo) Find() ([]*model.SysLogininfor, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysLogininfor), err
}

func (s sysLogininforDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysLogininfor, err error) {
	buf := make([]*model.SysLogininfor, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysLogininforDo) FindInBatches(result *[]*model.SysLogininfor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysLogininforDo) Attrs(attrs ...field.AssignExpr) *sysLogininforDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysLogininforDo) Assign(attrs ...field.AssignExpr) *sysLogininforDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysLogininforDo) Joins(fields ...field.RelationField) *sysLogininforDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysLogininforDo) Preload(fields ...field.RelationField) *sysLogininforDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysLogininforDo) FirstOrInit() (*model.SysLogininfor, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLogininfor), nil
	}
}

func (s sysLogininforDo) FirstOrCreate() (*model.SysLogininfor, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysLogininfor), nil
	}
}

func (s sysLogininforDo) FindByPage(offset int, limit int) (result []*model.SysLogininfor, count int64, err error) {
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

func (s sysLogininforDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysLogininforDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysLogininforDo) Delete(models ...*model.SysLogininfor) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysLogininforDo) withDO(do gen.Dao) *sysLogininforDo {
	s.DO = *do.(*gen.DO)
	return s
}
