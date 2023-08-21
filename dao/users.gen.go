// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"main/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUsers(db *gorm.DB, opts ...gen.DOOption) users {
	_users := users{}

	_users.usersDo.UseDB(db, opts...)
	_users.usersDo.UseModel(&model.Users{})

	tableName := _users.usersDo.TableName()
	_users.ALL = field.NewAsterisk(tableName)
	_users.UserID = field.NewInt(tableName, "user_id")
	_users.UserName = field.NewString(tableName, "user_name")
	_users.UserPassword = field.NewString(tableName, "user_password")

	_users.fillFieldMap()

	return _users
}

type users struct {
	usersDo

	ALL          field.Asterisk
	UserID       field.Int
	UserName     field.String
	UserPassword field.String

	fieldMap map[string]field.Expr
}

func (u users) Table(newTableName string) *users {
	u.usersDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u users) As(alias string) *users {
	u.usersDo.DO = *(u.usersDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *users) updateTableName(table string) *users {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewInt(table, "user_id")
	u.UserName = field.NewString(table, "user_name")
	u.UserPassword = field.NewString(table, "user_password")

	u.fillFieldMap()

	return u
}

func (u *users) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *users) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["user_password"] = u.UserPassword
}

func (u users) clone(db *gorm.DB) users {
	u.usersDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u users) replaceDB(db *gorm.DB) users {
	u.usersDo.ReplaceDB(db)
	return u
}

type usersDo struct{ gen.DO }

type IUsersDo interface {
	gen.SubQuery
	Debug() IUsersDo
	WithContext(ctx context.Context) IUsersDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUsersDo
	WriteDB() IUsersDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUsersDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUsersDo
	Not(conds ...gen.Condition) IUsersDo
	Or(conds ...gen.Condition) IUsersDo
	Select(conds ...field.Expr) IUsersDo
	Where(conds ...gen.Condition) IUsersDo
	Order(conds ...field.Expr) IUsersDo
	Distinct(cols ...field.Expr) IUsersDo
	Omit(cols ...field.Expr) IUsersDo
	Join(table schema.Tabler, on ...field.Expr) IUsersDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUsersDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUsersDo
	Group(cols ...field.Expr) IUsersDo
	Having(conds ...gen.Condition) IUsersDo
	Limit(limit int) IUsersDo
	Offset(offset int) IUsersDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersDo
	Unscoped() IUsersDo
	Create(values ...*model.Users) error
	CreateInBatches(values []*model.Users, batchSize int) error
	Save(values ...*model.Users) error
	First() (*model.Users, error)
	Take() (*model.Users, error)
	Last() (*model.Users, error)
	Find() ([]*model.Users, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Users, err error)
	FindInBatches(result *[]*model.Users, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Users) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUsersDo
	Assign(attrs ...field.AssignExpr) IUsersDo
	Joins(fields ...field.RelationField) IUsersDo
	Preload(fields ...field.RelationField) IUsersDo
	FirstOrInit() (*model.Users, error)
	FirstOrCreate() (*model.Users, error)
	FindByPage(offset int, limit int) (result []*model.Users, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUsersDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u usersDo) Debug() IUsersDo {
	return u.withDO(u.DO.Debug())
}

func (u usersDo) WithContext(ctx context.Context) IUsersDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u usersDo) ReadDB() IUsersDo {
	return u.Clauses(dbresolver.Read)
}

func (u usersDo) WriteDB() IUsersDo {
	return u.Clauses(dbresolver.Write)
}

func (u usersDo) Session(config *gorm.Session) IUsersDo {
	return u.withDO(u.DO.Session(config))
}

func (u usersDo) Clauses(conds ...clause.Expression) IUsersDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u usersDo) Returning(value interface{}, columns ...string) IUsersDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u usersDo) Not(conds ...gen.Condition) IUsersDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u usersDo) Or(conds ...gen.Condition) IUsersDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u usersDo) Select(conds ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u usersDo) Where(conds ...gen.Condition) IUsersDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u usersDo) Order(conds ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u usersDo) Distinct(cols ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u usersDo) Omit(cols ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u usersDo) Join(table schema.Tabler, on ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u usersDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUsersDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u usersDo) RightJoin(table schema.Tabler, on ...field.Expr) IUsersDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u usersDo) Group(cols ...field.Expr) IUsersDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u usersDo) Having(conds ...gen.Condition) IUsersDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u usersDo) Limit(limit int) IUsersDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u usersDo) Offset(offset int) IUsersDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u usersDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUsersDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u usersDo) Unscoped() IUsersDo {
	return u.withDO(u.DO.Unscoped())
}

func (u usersDo) Create(values ...*model.Users) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u usersDo) CreateInBatches(values []*model.Users, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u usersDo) Save(values ...*model.Users) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u usersDo) First() (*model.Users, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Users), nil
	}
}

func (u usersDo) Take() (*model.Users, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Users), nil
	}
}

func (u usersDo) Last() (*model.Users, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Users), nil
	}
}

func (u usersDo) Find() ([]*model.Users, error) {
	result, err := u.DO.Find()
	return result.([]*model.Users), err
}

func (u usersDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Users, err error) {
	buf := make([]*model.Users, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u usersDo) FindInBatches(result *[]*model.Users, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u usersDo) Attrs(attrs ...field.AssignExpr) IUsersDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u usersDo) Assign(attrs ...field.AssignExpr) IUsersDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u usersDo) Joins(fields ...field.RelationField) IUsersDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u usersDo) Preload(fields ...field.RelationField) IUsersDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u usersDo) FirstOrInit() (*model.Users, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Users), nil
	}
}

func (u usersDo) FirstOrCreate() (*model.Users, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Users), nil
	}
}

func (u usersDo) FindByPage(offset int, limit int) (result []*model.Users, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u usersDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u usersDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u usersDo) Delete(models ...*model.Users) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *usersDo) withDO(do gen.Dao) *usersDo {
	u.DO = *do.(*gen.DO)
	return u
}
