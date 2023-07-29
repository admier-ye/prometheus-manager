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

func newPromNodeDirFileGroupStrategy(db *gorm.DB, opts ...gen.DOOption) promNodeDirFileGroupStrategy {
	_promNodeDirFileGroupStrategy := promNodeDirFileGroupStrategy{}

	_promNodeDirFileGroupStrategy.promNodeDirFileGroupStrategyDo.UseDB(db, opts...)
	_promNodeDirFileGroupStrategy.promNodeDirFileGroupStrategyDo.UseModel(&model.PromNodeDirFileGroupStrategy{})

	tableName := _promNodeDirFileGroupStrategy.promNodeDirFileGroupStrategyDo.TableName()
	_promNodeDirFileGroupStrategy.ALL = field.NewAsterisk(tableName)
	_promNodeDirFileGroupStrategy.ID = field.NewInt32(tableName, "id")
	_promNodeDirFileGroupStrategy.GroupID = field.NewInt32(tableName, "group_id")
	_promNodeDirFileGroupStrategy.Alert = field.NewString(tableName, "alert")
	_promNodeDirFileGroupStrategy.Expr = field.NewString(tableName, "expr")
	_promNodeDirFileGroupStrategy.For = field.NewString(tableName, "for")
	_promNodeDirFileGroupStrategy.Labels = field.NewString(tableName, "labels")
	_promNodeDirFileGroupStrategy.Annotations = field.NewString(tableName, "annotations")
	_promNodeDirFileGroupStrategy.CreatedAt = field.NewTime(tableName, "created_at")
	_promNodeDirFileGroupStrategy.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promNodeDirFileGroupStrategy.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_promNodeDirFileGroupStrategy.fillFieldMap()

	return _promNodeDirFileGroupStrategy
}

type promNodeDirFileGroupStrategy struct {
	promNodeDirFileGroupStrategyDo

	ALL         field.Asterisk
	ID          field.Int32
	GroupID     field.Int32  // 策略组 ID
	Alert       field.String // 策略名称
	Expr        field.String // prom QL
	For         field.String // 持续时间
	Labels      field.String // 标签
	Annotations field.String // 内容
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	DeletedAt   field.Int64  // 删除时间

	fieldMap map[string]field.Expr
}

func (p promNodeDirFileGroupStrategy) Table(newTableName string) *promNodeDirFileGroupStrategy {
	p.promNodeDirFileGroupStrategyDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promNodeDirFileGroupStrategy) As(alias string) *promNodeDirFileGroupStrategy {
	p.promNodeDirFileGroupStrategyDo.DO = *(p.promNodeDirFileGroupStrategyDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promNodeDirFileGroupStrategy) updateTableName(table string) *promNodeDirFileGroupStrategy {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.GroupID = field.NewInt32(table, "group_id")
	p.Alert = field.NewString(table, "alert")
	p.Expr = field.NewString(table, "expr")
	p.For = field.NewString(table, "for")
	p.Labels = field.NewString(table, "labels")
	p.Annotations = field.NewString(table, "annotations")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewInt64(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *promNodeDirFileGroupStrategy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promNodeDirFileGroupStrategy) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["group_id"] = p.GroupID
	p.fieldMap["alert"] = p.Alert
	p.fieldMap["expr"] = p.Expr
	p.fieldMap["for"] = p.For
	p.fieldMap["labels"] = p.Labels
	p.fieldMap["annotations"] = p.Annotations
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p promNodeDirFileGroupStrategy) clone(db *gorm.DB) promNodeDirFileGroupStrategy {
	p.promNodeDirFileGroupStrategyDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promNodeDirFileGroupStrategy) replaceDB(db *gorm.DB) promNodeDirFileGroupStrategy {
	p.promNodeDirFileGroupStrategyDo.ReplaceDB(db)
	return p
}

type promNodeDirFileGroupStrategyDo struct{ gen.DO }

type IPromNodeDirFileGroupStrategyDo interface {
	gen.SubQuery
	Debug() IPromNodeDirFileGroupStrategyDo
	WithContext(ctx context.Context) IPromNodeDirFileGroupStrategyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromNodeDirFileGroupStrategyDo
	WriteDB() IPromNodeDirFileGroupStrategyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromNodeDirFileGroupStrategyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromNodeDirFileGroupStrategyDo
	Not(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo
	Or(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo
	Select(conds ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Where(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo
	Order(conds ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Distinct(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Omit(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Join(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Group(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo
	Having(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo
	Limit(limit int) IPromNodeDirFileGroupStrategyDo
	Offset(offset int) IPromNodeDirFileGroupStrategyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirFileGroupStrategyDo
	Unscoped() IPromNodeDirFileGroupStrategyDo
	Create(values ...*model.PromNodeDirFileGroupStrategy) error
	CreateInBatches(values []*model.PromNodeDirFileGroupStrategy, batchSize int) error
	Save(values ...*model.PromNodeDirFileGroupStrategy) error
	First() (*model.PromNodeDirFileGroupStrategy, error)
	Take() (*model.PromNodeDirFileGroupStrategy, error)
	Last() (*model.PromNodeDirFileGroupStrategy, error)
	Find() ([]*model.PromNodeDirFileGroupStrategy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDirFileGroupStrategy, err error)
	FindInBatches(result *[]*model.PromNodeDirFileGroupStrategy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromNodeDirFileGroupStrategy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromNodeDirFileGroupStrategyDo
	Assign(attrs ...field.AssignExpr) IPromNodeDirFileGroupStrategyDo
	Joins(fields ...field.RelationField) IPromNodeDirFileGroupStrategyDo
	Preload(fields ...field.RelationField) IPromNodeDirFileGroupStrategyDo
	FirstOrInit() (*model.PromNodeDirFileGroupStrategy, error)
	FirstOrCreate() (*model.PromNodeDirFileGroupStrategy, error)
	FindByPage(offset int, limit int) (result []*model.PromNodeDirFileGroupStrategy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromNodeDirFileGroupStrategyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	WhereID(ctx context.Context, id uint) (result *model.PromNodeDirFileGroupStrategy, err error)
}

// select * from @@table where id = @id
func (p promNodeDirFileGroupStrategyDo) WhereID(ctx context.Context, id uint) (result *model.PromNodeDirFileGroupStrategy, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_node_dir_file_group_strategies where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promNodeDirFileGroupStrategyDo) Debug() IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Debug())
}

func (p promNodeDirFileGroupStrategyDo) WithContext(ctx context.Context) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promNodeDirFileGroupStrategyDo) ReadDB() IPromNodeDirFileGroupStrategyDo {
	return p.Clauses(dbresolver.Read)
}

func (p promNodeDirFileGroupStrategyDo) WriteDB() IPromNodeDirFileGroupStrategyDo {
	return p.Clauses(dbresolver.Write)
}

func (p promNodeDirFileGroupStrategyDo) Session(config *gorm.Session) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Session(config))
}

func (p promNodeDirFileGroupStrategyDo) Clauses(conds ...clause.Expression) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Returning(value interface{}, columns ...string) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promNodeDirFileGroupStrategyDo) Not(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Or(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Select(conds ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Where(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Order(conds ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Distinct(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promNodeDirFileGroupStrategyDo) Omit(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promNodeDirFileGroupStrategyDo) Join(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promNodeDirFileGroupStrategyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promNodeDirFileGroupStrategyDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promNodeDirFileGroupStrategyDo) Group(cols ...field.Expr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promNodeDirFileGroupStrategyDo) Having(conds ...gen.Condition) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promNodeDirFileGroupStrategyDo) Limit(limit int) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promNodeDirFileGroupStrategyDo) Offset(offset int) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promNodeDirFileGroupStrategyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promNodeDirFileGroupStrategyDo) Unscoped() IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promNodeDirFileGroupStrategyDo) Create(values ...*model.PromNodeDirFileGroupStrategy) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promNodeDirFileGroupStrategyDo) CreateInBatches(values []*model.PromNodeDirFileGroupStrategy, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promNodeDirFileGroupStrategyDo) Save(values ...*model.PromNodeDirFileGroupStrategy) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promNodeDirFileGroupStrategyDo) First() (*model.PromNodeDirFileGroupStrategy, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroupStrategy), nil
	}
}

func (p promNodeDirFileGroupStrategyDo) Take() (*model.PromNodeDirFileGroupStrategy, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroupStrategy), nil
	}
}

func (p promNodeDirFileGroupStrategyDo) Last() (*model.PromNodeDirFileGroupStrategy, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroupStrategy), nil
	}
}

func (p promNodeDirFileGroupStrategyDo) Find() ([]*model.PromNodeDirFileGroupStrategy, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromNodeDirFileGroupStrategy), err
}

func (p promNodeDirFileGroupStrategyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromNodeDirFileGroupStrategy, err error) {
	buf := make([]*model.PromNodeDirFileGroupStrategy, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promNodeDirFileGroupStrategyDo) FindInBatches(result *[]*model.PromNodeDirFileGroupStrategy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promNodeDirFileGroupStrategyDo) Attrs(attrs ...field.AssignExpr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promNodeDirFileGroupStrategyDo) Assign(attrs ...field.AssignExpr) IPromNodeDirFileGroupStrategyDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promNodeDirFileGroupStrategyDo) Joins(fields ...field.RelationField) IPromNodeDirFileGroupStrategyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promNodeDirFileGroupStrategyDo) Preload(fields ...field.RelationField) IPromNodeDirFileGroupStrategyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promNodeDirFileGroupStrategyDo) FirstOrInit() (*model.PromNodeDirFileGroupStrategy, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroupStrategy), nil
	}
}

func (p promNodeDirFileGroupStrategyDo) FirstOrCreate() (*model.PromNodeDirFileGroupStrategy, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromNodeDirFileGroupStrategy), nil
	}
}

func (p promNodeDirFileGroupStrategyDo) FindByPage(offset int, limit int) (result []*model.PromNodeDirFileGroupStrategy, count int64, err error) {
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

func (p promNodeDirFileGroupStrategyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promNodeDirFileGroupStrategyDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promNodeDirFileGroupStrategyDo) Delete(models ...*model.PromNodeDirFileGroupStrategy) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promNodeDirFileGroupStrategyDo) withDO(do gen.Dao) *promNodeDirFileGroupStrategyDo {
	p.DO = *do.(*gen.DO)
	return p
}
