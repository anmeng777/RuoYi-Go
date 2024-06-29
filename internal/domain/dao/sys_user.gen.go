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

func newSysUser(db *gorm.DB, opts ...gen.DOOption) sysUser {
	_sysUser := sysUser{}

	_sysUser.sysUserDo.UseDB(db, opts...)
	_sysUser.sysUserDo.UseModel(&model.SysUser{})

	tableName := _sysUser.sysUserDo.TableName()
	_sysUser.ALL = field.NewAsterisk(tableName)
	_sysUser.UserID = field.NewInt64(tableName, "user_id")
	_sysUser.DeptID = field.NewInt64(tableName, "dept_id")
	_sysUser.UserName = field.NewString(tableName, "user_name")
	_sysUser.NickName = field.NewString(tableName, "nick_name")
	_sysUser.UserType = field.NewString(tableName, "user_type")
	_sysUser.Email = field.NewString(tableName, "email")
	_sysUser.Phonenumber = field.NewString(tableName, "phonenumber")
	_sysUser.Sex = field.NewString(tableName, "sex")
	_sysUser.Avatar = field.NewString(tableName, "avatar")
	_sysUser.Password = field.NewString(tableName, "password")
	_sysUser.Status = field.NewString(tableName, "status")
	_sysUser.DelFlag = field.NewString(tableName, "del_flag")
	_sysUser.LoginIP = field.NewString(tableName, "login_ip")
	_sysUser.LoginDate = field.NewTime(tableName, "login_date")
	_sysUser.CreateBy = field.NewString(tableName, "create_by")
	_sysUser.CreateTime = field.NewTime(tableName, "create_time")
	_sysUser.UpdateBy = field.NewString(tableName, "update_by")
	_sysUser.UpdateTime = field.NewTime(tableName, "update_time")
	_sysUser.Remark = field.NewString(tableName, "remark")

	_sysUser.fillFieldMap()

	return _sysUser
}

type sysUser struct {
	sysUserDo sysUserDo

	ALL         field.Asterisk
	UserID      field.Int64  // 用户ID
	DeptID      field.Int64  // 部门ID
	UserName    field.String // 用户账号
	NickName    field.String // 用户昵称
	UserType    field.String // 用户类型（00系统用户）
	Email       field.String // 用户邮箱
	Phonenumber field.String // 手机号码
	Sex         field.String // 用户性别（0男 1女 2未知）
	Avatar      field.String // 头像地址
	Password    field.String // 密码
	Status      field.String // 帐号状态（0正常 1停用）
	DelFlag     field.String // 删除标志（0代表存在 2代表删除）
	LoginIP     field.String // 最后登录IP
	LoginDate   field.Time   // 最后登录时间
	CreateBy    field.String // 创建者
	CreateTime  field.Time   // 创建时间
	UpdateBy    field.String // 更新者
	UpdateTime  field.Time   // 更新时间
	Remark      field.String // 备注

	fieldMap map[string]field.Expr
}

func (s sysUser) Table(newTableName string) *sysUser {
	s.sysUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUser) As(alias string) *sysUser {
	s.sysUserDo.DO = *(s.sysUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUser) updateTableName(table string) *sysUser {
	s.ALL = field.NewAsterisk(table)
	s.UserID = field.NewInt64(table, "user_id")
	s.DeptID = field.NewInt64(table, "dept_id")
	s.UserName = field.NewString(table, "user_name")
	s.NickName = field.NewString(table, "nick_name")
	s.UserType = field.NewString(table, "user_type")
	s.Email = field.NewString(table, "email")
	s.Phonenumber = field.NewString(table, "phonenumber")
	s.Sex = field.NewString(table, "sex")
	s.Avatar = field.NewString(table, "avatar")
	s.Password = field.NewString(table, "password")
	s.Status = field.NewString(table, "status")
	s.DelFlag = field.NewString(table, "del_flag")
	s.LoginIP = field.NewString(table, "login_ip")
	s.LoginDate = field.NewTime(table, "login_date")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")

	s.fillFieldMap()

	return s
}

func (s *sysUser) WithContext(ctx context.Context) *sysUserDo { return s.sysUserDo.WithContext(ctx) }

func (s sysUser) TableName() string { return s.sysUserDo.TableName() }

func (s sysUser) Alias() string { return s.sysUserDo.Alias() }

func (s sysUser) Columns(cols ...field.Expr) gen.Columns { return s.sysUserDo.Columns(cols...) }

func (s *sysUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 19)
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["dept_id"] = s.DeptID
	s.fieldMap["user_name"] = s.UserName
	s.fieldMap["nick_name"] = s.NickName
	s.fieldMap["user_type"] = s.UserType
	s.fieldMap["email"] = s.Email
	s.fieldMap["phonenumber"] = s.Phonenumber
	s.fieldMap["sex"] = s.Sex
	s.fieldMap["avatar"] = s.Avatar
	s.fieldMap["password"] = s.Password
	s.fieldMap["status"] = s.Status
	s.fieldMap["del_flag"] = s.DelFlag
	s.fieldMap["login_ip"] = s.LoginIP
	s.fieldMap["login_date"] = s.LoginDate
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
}

func (s sysUser) clone(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUser) replaceDB(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceDB(db)
	return s
}

type sysUserDo struct{ gen.DO }

func (s sysUserDo) Debug() *sysUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserDo) WithContext(ctx context.Context) *sysUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserDo) ReadDB() *sysUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserDo) WriteDB() *sysUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserDo) Session(config *gorm.Session) *sysUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserDo) Clauses(conds ...clause.Expression) *sysUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserDo) Returning(value interface{}, columns ...string) *sysUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserDo) Not(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserDo) Or(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserDo) Select(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserDo) Where(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserDo) Order(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserDo) Distinct(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserDo) Omit(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserDo) Join(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserDo) Group(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserDo) Having(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserDo) Limit(limit int) *sysUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserDo) Offset(offset int) *sysUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserDo) Unscoped() *sysUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserDo) Create(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserDo) CreateInBatches(values []*model.SysUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserDo) Save(values ...*model.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserDo) First() (*model.SysUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Take() (*model.SysUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Last() (*model.SysUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) Find() ([]*model.SysUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysUser), err
}

func (s sysUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysUser, err error) {
	buf := make([]*model.SysUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserDo) FindInBatches(result *[]*model.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserDo) Attrs(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserDo) Assign(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserDo) Joins(fields ...field.RelationField) *sysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserDo) Preload(fields ...field.RelationField) *sysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserDo) FirstOrInit() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FirstOrCreate() (*model.SysUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysUser), nil
	}
}

func (s sysUserDo) FindByPage(offset int, limit int) (result []*model.SysUser, count int64, err error) {
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

func (s sysUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserDo) Delete(models ...*model.SysUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserDo) withDO(do gen.Dao) *sysUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
