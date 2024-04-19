package contextstore

import "gorm.io/gorm"

type (
	ConditionBuilder interface {
		Find(dest any, conds ...any) ConditionBuilder
		Select(query any, args ...any) ConditionBuilder
		Distinct(args ...any) ConditionBuilder
		Table(name string, args ...any) ConditionBuilder
		Not(query any, args ...any) ConditionBuilder
		Or(query any, args ...any) ConditionBuilder
		Where(query any, args ...any) ConditionBuilder
		Omit(columns ...string) ConditionBuilder
		Joins(query string, args ...any) ConditionBuilder
		InnerJoins(query string, args ...any) ConditionBuilder
		Group(name string) ConditionBuilder
		Order(value any) ConditionBuilder
		Limit(limit int) ConditionBuilder
		Offset(offset int) ConditionBuilder
		Having(query any, args ...any) ConditionBuilder
		Raw(sql string, values ...any) ConditionBuilder
	}
	builder struct {
		db *gorm.DB
	}
)

func NewConditionBuilder(db *gorm.DB) ConditionBuilder {
	return &builder{
		db: db,
	}
}

func (ctx *builder) Find(dest any, conds ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Find(dest, conds))
}

func (ctx *builder) Select(query any, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Select(query, args))
}

func (ctx *builder) Distinct(args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Distinct(args))
}

func (ctx *builder) Table(name string, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Table(name, args))
}

func (ctx *builder) Not(query any, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Not(query, args))
}

func (ctx *builder) Or(query any, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Or(query, args))
}

func (ctx *builder) Where(query any, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Where(query, args))
}

func (ctx *builder) Omit(columns ...string) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Omit(columns...))
}

func (ctx *builder) Joins(query string, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Joins(query, args))
}

func (ctx *builder) InnerJoins(query string, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.InnerJoins(query, args))
}

func (ctx *builder) Group(name string) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Group(name))
}

func (ctx *builder) Order(value any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Order(value))
}

func (ctx *builder) Limit(limit int) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Limit(limit))
}

func (ctx *builder) Offset(offset int) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Offset(offset))
}

func (ctx *builder) Having(query any, args ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Having(query, args))
}

func (ctx *builder) Raw(sql string, values ...any) ConditionBuilder {
	return NewConditionBuilder(ctx.db.Raw(sql, values...))
}
