// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"prometheus-manager/dal/model"
)

func newPromNodeDir(db *gorm.DB, opts ...gen.DOOption) promNodeDir {
	_promNodeDir := promNodeDir{}

	_promNodeDir.promNodeDirDo.UseDB(db, opts...)
	_promNodeDir.promNodeDirDo.UseModel(&model.PromNodeDir{})

	tableName := _promNodeDir.promNodeDirDo.TableName()
	_promNodeDir.ALL = field.NewAsterisk(tableName)
	_promNodeDir.ID = field.NewInt32(tableName, "id")
	_promNodeDir.NodeID = field.NewInt32(tableName, "node_id")
	_promNodeDir.Path = field.NewString(tableName, "path")
	_promNodeDir.CreatedAt = field.NewTime(tableName, "created_at")
	_promNodeDir.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promNodeDir.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_promNodeDir.fillFieldMap()

	return _promNodeDir
}

type promNodeDir struct {
	promNodeDirDo

	ALL       field.Asterisk
	ID        field.Int32
	NodeID    field.Int32  // 节点 ID
	Path      field.String // 目录地址
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (p promNodeDir) Table(newTableName string) *promNodeDir {
	p.promNodeDirDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promNodeDir) As(alias string) *promNodeDir {
	p.promNodeDirDo.DO = *(p.promNodeDirDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promNodeDir) updateTableName(table string) *promNodeDir {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.NodeID = field.NewInt32(table, "node_id")
	p.Path = field.NewString(table, "path")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewInt64(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *promNodeDir) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promNodeDir) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["node_id"] = p.NodeID
	p.fieldMap["path"] = p.Path
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p promNodeDir) clone(db *gorm.DB) promNodeDir {
	p.promNodeDirDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promNodeDir) replaceDB(db *gorm.DB) promNodeDir {
	p.promNodeDirDo.ReplaceDB(db)
	return p
}

type promNodeDirDo struct{ gen.DO }

type IPromNodeDirDo interface {
	gen.SubQuery
	Debug() IPromNodeDirDo
	WithContext(ctx context.Context) IPromNodeDirDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromNodeDirDo
	WriteDB() IPromNodeDirDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromNodeDirDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromNodeDirDo
	Not(conds ...gen.Condition) IPromNodeDirDo
	Or(conds ...gen.Condition) IPromNodeDirDo
	Select(conds ...field.Expr) IPromNodeDirDo
	Where(conds ...gen.Condition) IPromNodeDirDo
	Order(conds ...field.Expr) IPromNodeDirDo
	Distinct(cols ...field.Expr) IPromNodeDirDo
	Omit(cols ...field.Expr) IPromNodeDirDo
	Join(table schema.Tabler, on ...field.Expr) IPromNodeDirDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirDo
	Group(cols ...field.Expr) IPromNodeDirDo
	Having(conds ...gen.Condition) IPromNodeDirDo
	Limit(limit int) IPromNodeDirDo
	Offset(offset int) IPromNodeDirDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirDo
	Unscoped() IPromNodeDirDo
	Create(values ...*model.PromNodeDir) error
	CreateInBatches(values []*model.PromNodeDir, batchSize int) error
	Save(values ...*model.PromNodeDir) error
	First() (*model.PromNodeDir, error)
	Take() (*model.PromNodeDir, error)
	Last() (*model.PromNodeDir, error)
	Find() ([]*model.PromNodeDir, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDir, err error)
	FindInBatches(result *[]*model.PromNodeDir, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromNodeDir) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromNodeDirDo
	Assign(attrs ...field.AssignExpr) IPromNodeDirDo
	Joins(fields ...field.RelationField) IPromNodeDirDo
	Preload(fields ...field.RelationField) IPromNodeDirDo
	FirstOrInit() (*model.PromNodeDir, error)
	FirstOrCreate() (*model.PromNodeDir, error)
	FindByPage(offset int, limit int) (result []*model.PromNodeDir, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromNodeDirDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	WhereID(ctx context.Context, id uint) (result *model.PromNodeDir, err error)
}

// select * from @@table where id = @id
func (p promNodeDirDo) WhereID(ctx context.Context, id uint) (result *model.PromNodeDir, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_node_dirs where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promNodeDirDo) Debug() IPromNodeDirDo {
	return p.withDO(p.DO.Debug())
}

func (p promNodeDirDo) WithContext(ctx context.Context) IPromNodeDirDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promNodeDirDo) ReadDB() IPromNodeDirDo {
	return p.Clauses(dbresolver.Read)
}

func (p promNodeDirDo) WriteDB() IPromNodeDirDo {
	return p.Clauses(dbresolver.Write)
}

func (p promNodeDirDo) Session(config *gorm.Session) IPromNodeDirDo {
	return p.withDO(p.DO.Session(config))
}

func (p promNodeDirDo) Clauses(conds ...clause.Expression) IPromNodeDirDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promNodeDirDo) Returning(value interface{}, columns ...string) IPromNodeDirDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promNodeDirDo) Not(conds ...gen.Condition) IPromNodeDirDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promNodeDirDo) Or(conds ...gen.Condition) IPromNodeDirDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promNodeDirDo) Select(conds ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promNodeDirDo) Where(conds ...gen.Condition) IPromNodeDirDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promNodeDirDo) Order(conds ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promNodeDirDo) Distinct(cols ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promNodeDirDo) Omit(cols ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promNodeDirDo) Join(table schema.Tabler, on ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promNodeDirDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promNodeDirDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promNodeDirDo) Group(cols ...field.Expr) IPromNodeDirDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promNodeDirDo) Having(conds ...gen.Condition) IPromNodeDirDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promNodeDirDo) Limit(limit int) IPromNodeDirDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promNodeDirDo) Offset(offset int) IPromNodeDirDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promNodeDirDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promNodeDirDo) Unscoped() IPromNodeDirDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promNodeDirDo) Create(values ...*model.PromNodeDir) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promNodeDirDo) CreateInBatches(values []*model.PromNodeDir, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promNodeDirDo) Save(values ...*model.PromNodeDir) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promNodeDirDo) First() (*model.PromNodeDir, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDir), nil
	}
}

func (p promNodeDirDo) Take() (*model.PromNodeDir, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDir), nil
	}
}

func (p promNodeDirDo) Last() (*model.PromNodeDir, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDir), nil
	}
}

func (p promNodeDirDo) Find() ([]*model.PromNodeDir, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromNodeDir), err
}

func (p promNodeDirDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDir, err error) {
	buf := make([]*model.PromNodeDir, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promNodeDirDo) FindInBatches(result *[]*model.PromNodeDir, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promNodeDirDo) Attrs(attrs ...field.AssignExpr) IPromNodeDirDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promNodeDirDo) Assign(attrs ...field.AssignExpr) IPromNodeDirDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promNodeDirDo) Joins(fields ...field.RelationField) IPromNodeDirDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promNodeDirDo) Preload(fields ...field.RelationField) IPromNodeDirDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promNodeDirDo) FirstOrInit() (*model.PromNodeDir, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDir), nil
	}
}

func (p promNodeDirDo) FirstOrCreate() (*model.PromNodeDir, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDir), nil
	}
}

func (p promNodeDirDo) FindByPage(offset int, limit int) (result []*model.PromNodeDir, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p promNodeDirDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promNodeDirDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promNodeDirDo) Delete(models ...*model.PromNodeDir) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promNodeDirDo) withDO(do gen.Dao) *promNodeDirDo {
	p.DO = *do.(*gen.DO)
	return p
}
