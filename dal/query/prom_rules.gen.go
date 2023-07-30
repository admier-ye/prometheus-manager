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

func newPromRule(db *gorm.DB, opts ...gen.DOOption) promRule {
	_promRule := promRule{}

	_promRule.promRuleDo.UseDB(db, opts...)
	_promRule.promRuleDo.UseModel(&model.PromRule{})

	tableName := _promRule.promRuleDo.TableName()
	_promRule.ALL = field.NewAsterisk(tableName)
	_promRule.ID = field.NewInt32(tableName, "id")
	_promRule.Alert = field.NewString(tableName, "alert")
	_promRule.Expr = field.NewString(tableName, "expr")
	_promRule.For = field.NewString(tableName, "for")
	_promRule.Labels = field.NewString(tableName, "labels")
	_promRule.Annotations = field.NewString(tableName, "annotations")
	_promRule.CreatedAt = field.NewTime(tableName, "created_at")
	_promRule.UpdatedAt = field.NewTime(tableName, "updated_at")
	_promRule.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_promRule.fillFieldMap()

	return _promRule
}

type promRule struct {
	promRuleDo

	ALL         field.Asterisk
	ID          field.Int32
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

func (p promRule) Table(newTableName string) *promRule {
	p.promRuleDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p promRule) As(alias string) *promRule {
	p.promRuleDo.DO = *(p.promRuleDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *promRule) updateTableName(table string) *promRule {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
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

func (p *promRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *promRule) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["alert"] = p.Alert
	p.fieldMap["expr"] = p.Expr
	p.fieldMap["for"] = p.For
	p.fieldMap["labels"] = p.Labels
	p.fieldMap["annotations"] = p.Annotations
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p promRule) clone(db *gorm.DB) promRule {
	p.promRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p promRule) replaceDB(db *gorm.DB) promRule {
	p.promRuleDo.ReplaceDB(db)
	return p
}

type promRuleDo struct{ gen.DO }

type IPromRuleDo interface {
	gen.SubQuery
	Debug() IPromRuleDo
	WithContext(ctx context.Context) IPromRuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPromRuleDo
	WriteDB() IPromRuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPromRuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPromRuleDo
	Not(conds ...gen.Condition) IPromRuleDo
	Or(conds ...gen.Condition) IPromRuleDo
	Select(conds ...field.Expr) IPromRuleDo
	Where(conds ...gen.Condition) IPromRuleDo
	Order(conds ...field.Expr) IPromRuleDo
	Distinct(cols ...field.Expr) IPromRuleDo
	Omit(cols ...field.Expr) IPromRuleDo
	Join(table schema.Tabler, on ...field.Expr) IPromRuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPromRuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPromRuleDo
	Group(cols ...field.Expr) IPromRuleDo
	Having(conds ...gen.Condition) IPromRuleDo
	Limit(limit int) IPromRuleDo
	Offset(offset int) IPromRuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPromRuleDo
	Unscoped() IPromRuleDo
	Create(values ...*model.PromRule) error
	CreateInBatches(values []*model.PromRule, batchSize int) error
	Save(values ...*model.PromRule) error
	First() (*model.PromRule, error)
	Take() (*model.PromRule, error)
	Last() (*model.PromRule, error)
	Find() ([]*model.PromRule, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromRule, err error)
	FindInBatches(result *[]*model.PromRule, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PromRule) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPromRuleDo
	Assign(attrs ...field.AssignExpr) IPromRuleDo
	Joins(fields ...field.RelationField) IPromRuleDo
	Preload(fields ...field.RelationField) IPromRuleDo
	FirstOrInit() (*model.PromRule, error)
	FirstOrCreate() (*model.PromRule, error)
	FindByPage(offset int, limit int) (result []*model.PromRule, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPromRuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindById(ctx context.Context, id int32) (result *model.PromRule, err error)
}

// select * from @@table where id = @id
func (p promRuleDo) FindById(ctx context.Context, id int32) (result *model.PromRule, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("select * from prom_rules where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p promRuleDo) Debug() IPromRuleDo {
	return p.withDO(p.DO.Debug())
}

func (p promRuleDo) WithContext(ctx context.Context) IPromRuleDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p promRuleDo) ReadDB() IPromRuleDo {
	return p.Clauses(dbresolver.Read)
}

func (p promRuleDo) WriteDB() IPromRuleDo {
	return p.Clauses(dbresolver.Write)
}

func (p promRuleDo) Session(config *gorm.Session) IPromRuleDo {
	return p.withDO(p.DO.Session(config))
}

func (p promRuleDo) Clauses(conds ...clause.Expression) IPromRuleDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p promRuleDo) Returning(value interface{}, columns ...string) IPromRuleDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p promRuleDo) Not(conds ...gen.Condition) IPromRuleDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p promRuleDo) Or(conds ...gen.Condition) IPromRuleDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p promRuleDo) Select(conds ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p promRuleDo) Where(conds ...gen.Condition) IPromRuleDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p promRuleDo) Order(conds ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p promRuleDo) Distinct(cols ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p promRuleDo) Omit(cols ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p promRuleDo) Join(table schema.Tabler, on ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p promRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p promRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p promRuleDo) Group(cols ...field.Expr) IPromRuleDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p promRuleDo) Having(conds ...gen.Condition) IPromRuleDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p promRuleDo) Limit(limit int) IPromRuleDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p promRuleDo) Offset(offset int) IPromRuleDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p promRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPromRuleDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p promRuleDo) Unscoped() IPromRuleDo {
	return p.withDO(p.DO.Unscoped())
}

func (p promRuleDo) Create(values ...*model.PromRule) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p promRuleDo) CreateInBatches(values []*model.PromRule, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p promRuleDo) Save(values ...*model.PromRule) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p promRuleDo) First() (*model.PromRule, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromRule), nil
	}
}

func (p promRuleDo) Take() (*model.PromRule, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromRule), nil
	}
}

func (p promRuleDo) Last() (*model.PromRule, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromRule), nil
	}
}

func (p promRuleDo) Find() ([]*model.PromRule, error) {
	result, err := p.DO.Find()
	return result.([]*model.PromRule), err
}

func (p promRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PromRule, err error) {
	buf := make([]*model.PromRule, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p promRuleDo) FindInBatches(result *[]*model.PromRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p promRuleDo) Attrs(attrs ...field.AssignExpr) IPromRuleDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p promRuleDo) Assign(attrs ...field.AssignExpr) IPromRuleDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p promRuleDo) Joins(fields ...field.RelationField) IPromRuleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p promRuleDo) Preload(fields ...field.RelationField) IPromRuleDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p promRuleDo) FirstOrInit() (*model.PromRule, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromRule), nil
	}
}

func (p promRuleDo) FirstOrCreate() (*model.PromRule, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PromRule), nil
	}
}

func (p promRuleDo) FindByPage(offset int, limit int) (result []*model.PromRule, count int64, err error) {
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

func (p promRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p promRuleDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p promRuleDo) Delete(models ...*model.PromRule) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *promRuleDo) withDO(do gen.Dao) *promRuleDo {
	p.DO = *do.(*gen.DO)
	return p
}
