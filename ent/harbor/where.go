// Code generated by entc, DO NOT EDIT.

package harbor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/andrewmthomas87/catan/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// X applies equality check predicate on the "x" field. It's identical to XEQ.
func X(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldX), v))
	})
}

// Y applies equality check predicate on the "y" field. It's identical to YEQ.
func Y(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldY), v))
	})
}

// XEQ applies the EQ predicate on the "x" field.
func XEQ(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldX), v))
	})
}

// XNEQ applies the NEQ predicate on the "x" field.
func XNEQ(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldX), v))
	})
}

// XIn applies the In predicate on the "x" field.
func XIn(vs ...int) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldX), v...))
	})
}

// XNotIn applies the NotIn predicate on the "x" field.
func XNotIn(vs ...int) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldX), v...))
	})
}

// XGT applies the GT predicate on the "x" field.
func XGT(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldX), v))
	})
}

// XGTE applies the GTE predicate on the "x" field.
func XGTE(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldX), v))
	})
}

// XLT applies the LT predicate on the "x" field.
func XLT(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldX), v))
	})
}

// XLTE applies the LTE predicate on the "x" field.
func XLTE(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldX), v))
	})
}

// YEQ applies the EQ predicate on the "y" field.
func YEQ(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldY), v))
	})
}

// YNEQ applies the NEQ predicate on the "y" field.
func YNEQ(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldY), v))
	})
}

// YIn applies the In predicate on the "y" field.
func YIn(vs ...int) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldY), v...))
	})
}

// YNotIn applies the NotIn predicate on the "y" field.
func YNotIn(vs ...int) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldY), v...))
	})
}

// YGT applies the GT predicate on the "y" field.
func YGT(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldY), v))
	})
}

// YGTE applies the GTE predicate on the "y" field.
func YGTE(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldY), v))
	})
}

// YLT applies the LT predicate on the "y" field.
func YLT(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldY), v))
	})
}

// YLTE applies the LTE predicate on the "y" field.
func YLTE(v int) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldY), v))
	})
}

// ResourceEQ applies the EQ predicate on the "resource" field.
func ResourceEQ(v Resource) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResource), v))
	})
}

// ResourceNEQ applies the NEQ predicate on the "resource" field.
func ResourceNEQ(v Resource) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResource), v))
	})
}

// ResourceIn applies the In predicate on the "resource" field.
func ResourceIn(vs ...Resource) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResource), v...))
	})
}

// ResourceNotIn applies the NotIn predicate on the "resource" field.
func ResourceNotIn(vs ...Resource) predicate.Harbor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Harbor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResource), v...))
	})
}

// HasSettlement applies the HasEdge predicate on the "settlement" edge.
func HasSettlement() predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SettlementTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, SettlementTable, SettlementColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSettlementWith applies the HasEdge predicate on the "settlement" edge with a given conditions (other predicates).
func HasSettlementWith(preds ...predicate.Settlement) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SettlementInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, SettlementTable, SettlementColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Harbor) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Harbor) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
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
func Not(p predicate.Harbor) predicate.Harbor {
	return predicate.Harbor(func(s *sql.Selector) {
		p(s.Not())
	})
}
