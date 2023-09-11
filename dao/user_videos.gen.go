// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"main/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUserVideo(db *gorm.DB, opts ...gen.DOOption) userVideo {
	_userVideo := userVideo{}

	_userVideo.userVideoDo.UseDB(db, opts...)
	_userVideo.userVideoDo.UseModel(&models.UserVideo{})

	tableName := _userVideo.userVideoDo.TableName()
	_userVideo.ALL = field.NewAsterisk(tableName)
	_userVideo.UserID = field.NewUint(tableName, "user_id")
	_userVideo.VideoID = field.NewUint(tableName, "video_id")
	_userVideo.User = userVideoBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "models.User"),
	}

	_userVideo.Video = userVideoBelongsToVideo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Video", "models.Video"),
	}

	_userVideo.fillFieldMap()

	return _userVideo
}

type userVideo struct {
	userVideoDo

	ALL     field.Asterisk
	UserID  field.Uint
	VideoID field.Uint
	User    userVideoBelongsToUser

	Video userVideoBelongsToVideo

	fieldMap map[string]field.Expr
}

func (u userVideo) Table(newTableName string) *userVideo {
	u.userVideoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userVideo) As(alias string) *userVideo {
	u.userVideoDo.DO = *(u.userVideoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userVideo) updateTableName(table string) *userVideo {
	u.ALL = field.NewAsterisk(table)
	u.UserID = field.NewUint(table, "user_id")
	u.VideoID = field.NewUint(table, "video_id")

	u.fillFieldMap()

	return u
}

func (u *userVideo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userVideo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["video_id"] = u.VideoID

}

func (u userVideo) clone(db *gorm.DB) userVideo {
	u.userVideoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userVideo) replaceDB(db *gorm.DB) userVideo {
	u.userVideoDo.ReplaceDB(db)
	return u
}

type userVideoBelongsToUser struct {
	db *gorm.DB

	field.RelationField
}

func (a userVideoBelongsToUser) Where(conds ...field.Expr) *userVideoBelongsToUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userVideoBelongsToUser) WithContext(ctx context.Context) *userVideoBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userVideoBelongsToUser) Session(session *gorm.Session) *userVideoBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a userVideoBelongsToUser) Model(m *models.UserVideo) *userVideoBelongsToUserTx {
	return &userVideoBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type userVideoBelongsToUserTx struct{ tx *gorm.Association }

func (a userVideoBelongsToUserTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userVideoBelongsToUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userVideoBelongsToUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userVideoBelongsToUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userVideoBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a userVideoBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type userVideoBelongsToVideo struct {
	db *gorm.DB

	field.RelationField
}

func (a userVideoBelongsToVideo) Where(conds ...field.Expr) *userVideoBelongsToVideo {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userVideoBelongsToVideo) WithContext(ctx context.Context) *userVideoBelongsToVideo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userVideoBelongsToVideo) Session(session *gorm.Session) *userVideoBelongsToVideo {
	a.db = a.db.Session(session)
	return &a
}

func (a userVideoBelongsToVideo) Model(m *models.UserVideo) *userVideoBelongsToVideoTx {
	return &userVideoBelongsToVideoTx{a.db.Model(m).Association(a.Name())}
}

type userVideoBelongsToVideoTx struct{ tx *gorm.Association }

func (a userVideoBelongsToVideoTx) Find() (result *models.Video, err error) {
	return result, a.tx.Find(&result)
}

func (a userVideoBelongsToVideoTx) Append(values ...*models.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userVideoBelongsToVideoTx) Replace(values ...*models.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userVideoBelongsToVideoTx) Delete(values ...*models.Video) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userVideoBelongsToVideoTx) Clear() error {
	return a.tx.Clear()
}

func (a userVideoBelongsToVideoTx) Count() int64 {
	return a.tx.Count()
}

type userVideoDo struct{ gen.DO }

type IUserVideoDo interface {
	gen.SubQuery
	Debug() IUserVideoDo
	WithContext(ctx context.Context) IUserVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserVideoDo
	WriteDB() IUserVideoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserVideoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserVideoDo
	Not(conds ...gen.Condition) IUserVideoDo
	Or(conds ...gen.Condition) IUserVideoDo
	Select(conds ...field.Expr) IUserVideoDo
	Where(conds ...gen.Condition) IUserVideoDo
	Order(conds ...field.Expr) IUserVideoDo
	Distinct(cols ...field.Expr) IUserVideoDo
	Omit(cols ...field.Expr) IUserVideoDo
	Join(table schema.Tabler, on ...field.Expr) IUserVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo
	Group(cols ...field.Expr) IUserVideoDo
	Having(conds ...gen.Condition) IUserVideoDo
	Limit(limit int) IUserVideoDo
	Offset(offset int) IUserVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVideoDo
	Unscoped() IUserVideoDo
	Create(values ...*models.UserVideo) error
	CreateInBatches(values []*models.UserVideo, batchSize int) error
	Save(values ...*models.UserVideo) error
	First() (*models.UserVideo, error)
	Take() (*models.UserVideo, error)
	Last() (*models.UserVideo, error)
	Find() ([]*models.UserVideo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserVideo, err error)
	FindInBatches(result *[]*models.UserVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.UserVideo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserVideoDo
	Assign(attrs ...field.AssignExpr) IUserVideoDo
	Joins(fields ...field.RelationField) IUserVideoDo
	Preload(fields ...field.RelationField) IUserVideoDo
	FirstOrInit() (*models.UserVideo, error)
	FirstOrCreate() (*models.UserVideo, error)
	FindByPage(offset int, limit int) (result []*models.UserVideo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userVideoDo) Debug() IUserVideoDo {
	return u.withDO(u.DO.Debug())
}

func (u userVideoDo) WithContext(ctx context.Context) IUserVideoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userVideoDo) ReadDB() IUserVideoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userVideoDo) WriteDB() IUserVideoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userVideoDo) Session(config *gorm.Session) IUserVideoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userVideoDo) Clauses(conds ...clause.Expression) IUserVideoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userVideoDo) Returning(value interface{}, columns ...string) IUserVideoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userVideoDo) Not(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userVideoDo) Or(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userVideoDo) Select(conds ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userVideoDo) Where(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userVideoDo) Order(conds ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userVideoDo) Distinct(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userVideoDo) Omit(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userVideoDo) Join(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userVideoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userVideoDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userVideoDo) Group(cols ...field.Expr) IUserVideoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userVideoDo) Having(conds ...gen.Condition) IUserVideoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userVideoDo) Limit(limit int) IUserVideoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userVideoDo) Offset(offset int) IUserVideoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userVideoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVideoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userVideoDo) Unscoped() IUserVideoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userVideoDo) Create(values ...*models.UserVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userVideoDo) CreateInBatches(values []*models.UserVideo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userVideoDo) Save(values ...*models.UserVideo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userVideoDo) First() (*models.UserVideo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserVideo), nil
	}
}

func (u userVideoDo) Take() (*models.UserVideo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserVideo), nil
	}
}

func (u userVideoDo) Last() (*models.UserVideo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserVideo), nil
	}
}

func (u userVideoDo) Find() ([]*models.UserVideo, error) {
	result, err := u.DO.Find()
	return result.([]*models.UserVideo), err
}

func (u userVideoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserVideo, err error) {
	buf := make([]*models.UserVideo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userVideoDo) FindInBatches(result *[]*models.UserVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userVideoDo) Attrs(attrs ...field.AssignExpr) IUserVideoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userVideoDo) Assign(attrs ...field.AssignExpr) IUserVideoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userVideoDo) Joins(fields ...field.RelationField) IUserVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userVideoDo) Preload(fields ...field.RelationField) IUserVideoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userVideoDo) FirstOrInit() (*models.UserVideo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserVideo), nil
	}
}

func (u userVideoDo) FirstOrCreate() (*models.UserVideo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserVideo), nil
	}
}

func (u userVideoDo) FindByPage(offset int, limit int) (result []*models.UserVideo, count int64, err error) {
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

func (u userVideoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userVideoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userVideoDo) Delete(models ...*models.UserVideo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userVideoDo) withDO(do gen.Dao) *userVideoDo {
	u.DO = *do.(*gen.DO)
	return u
}
