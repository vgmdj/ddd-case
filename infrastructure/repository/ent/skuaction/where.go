// Code generated by entc, DO NOT EDIT.

package skuaction

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// SkuName applies equality check predicate on the "sku_name" field. It's identical to SkuNameEQ.
func SkuName(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuName), v))
	})
}

// SkuNameCn applies equality check predicate on the "sku_name_cn" field. It's identical to SkuNameCnEQ.
func SkuNameCn(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuNameCn), v))
	})
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppID), v))
	})
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppID), v))
	})
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppID), v))
	})
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppID), v))
	})
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppID), v))
	})
}

// SkuNameEQ applies the EQ predicate on the "sku_name" field.
func SkuNameEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuName), v))
	})
}

// SkuNameNEQ applies the NEQ predicate on the "sku_name" field.
func SkuNameNEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuName), v))
	})
}

// SkuNameIn applies the In predicate on the "sku_name" field.
func SkuNameIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuName), v...))
	})
}

// SkuNameNotIn applies the NotIn predicate on the "sku_name" field.
func SkuNameNotIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuName), v...))
	})
}

// SkuNameGT applies the GT predicate on the "sku_name" field.
func SkuNameGT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuName), v))
	})
}

// SkuNameGTE applies the GTE predicate on the "sku_name" field.
func SkuNameGTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuName), v))
	})
}

// SkuNameLT applies the LT predicate on the "sku_name" field.
func SkuNameLT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuName), v))
	})
}

// SkuNameLTE applies the LTE predicate on the "sku_name" field.
func SkuNameLTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuName), v))
	})
}

// SkuNameContains applies the Contains predicate on the "sku_name" field.
func SkuNameContains(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSkuName), v))
	})
}

// SkuNameHasPrefix applies the HasPrefix predicate on the "sku_name" field.
func SkuNameHasPrefix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSkuName), v))
	})
}

// SkuNameHasSuffix applies the HasSuffix predicate on the "sku_name" field.
func SkuNameHasSuffix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSkuName), v))
	})
}

// SkuNameEqualFold applies the EqualFold predicate on the "sku_name" field.
func SkuNameEqualFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSkuName), v))
	})
}

// SkuNameContainsFold applies the ContainsFold predicate on the "sku_name" field.
func SkuNameContainsFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSkuName), v))
	})
}

// SkuNameCnEQ applies the EQ predicate on the "sku_name_cn" field.
func SkuNameCnEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnNEQ applies the NEQ predicate on the "sku_name_cn" field.
func SkuNameCnNEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnIn applies the In predicate on the "sku_name_cn" field.
func SkuNameCnIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkuNameCn), v...))
	})
}

// SkuNameCnNotIn applies the NotIn predicate on the "sku_name_cn" field.
func SkuNameCnNotIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkuNameCn), v...))
	})
}

// SkuNameCnGT applies the GT predicate on the "sku_name_cn" field.
func SkuNameCnGT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnGTE applies the GTE predicate on the "sku_name_cn" field.
func SkuNameCnGTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnLT applies the LT predicate on the "sku_name_cn" field.
func SkuNameCnLT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnLTE applies the LTE predicate on the "sku_name_cn" field.
func SkuNameCnLTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnContains applies the Contains predicate on the "sku_name_cn" field.
func SkuNameCnContains(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnHasPrefix applies the HasPrefix predicate on the "sku_name_cn" field.
func SkuNameCnHasPrefix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnHasSuffix applies the HasSuffix predicate on the "sku_name_cn" field.
func SkuNameCnHasSuffix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnEqualFold applies the EqualFold predicate on the "sku_name_cn" field.
func SkuNameCnEqualFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSkuNameCn), v))
	})
}

// SkuNameCnContainsFold applies the ContainsFold predicate on the "sku_name_cn" field.
func SkuNameCnContainsFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSkuNameCn), v))
	})
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAction), v))
	})
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAction), v...))
	})
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...string) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAction), v...))
	})
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAction), v))
	})
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAction), v))
	})
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAction), v))
	})
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAction), v))
	})
}

// ActionContains applies the Contains predicate on the "action" field.
func ActionContains(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAction), v))
	})
}

// ActionHasPrefix applies the HasPrefix predicate on the "action" field.
func ActionHasPrefix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAction), v))
	})
}

// ActionHasSuffix applies the HasSuffix predicate on the "action" field.
func ActionHasSuffix(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAction), v))
	})
}

// ActionEqualFold applies the EqualFold predicate on the "action" field.
func ActionEqualFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAction), v))
	})
}

// ActionContainsFold applies the ContainsFold predicate on the "action" field.
func ActionContainsFold(v string) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAction), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.SkuAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SkuAction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// HasConfig applies the HasEdge predicate on the "config" edge.
func HasConfig() predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigWith applies the HasEdge predicate on the "config" edge with a given conditions (other predicates).
func HasConfigWith(preds ...predicate.SkuConfig) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTemplate applies the HasEdge predicate on the "template" edge.
func HasTemplate() predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplateTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TemplateTable, TemplateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTemplateWith applies the HasEdge predicate on the "template" edge with a given conditions (other predicates).
func HasTemplateWith(preds ...predicate.SkuTemplate) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TemplateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TemplateTable, TemplateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusCode applies the HasEdge predicate on the "status_code" edge.
func HasStatusCode() predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusCodeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusCodeTable, StatusCodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusCodeWith applies the HasEdge predicate on the "status_code" edge with a given conditions (other predicates).
func HasStatusCodeWith(preds ...predicate.SkuStatusCode) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusCodeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusCodeTable, StatusCodeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.SkuAction) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.SkuAction) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SkuAction) predicate.SkuAction {
	return predicate.SkuAction(func(s *sql.Selector) {
		p(s.Not())
	})
}
