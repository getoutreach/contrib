// Code generated by entc, DO NOT EDIT.

package user

import (
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// Unnecessary applies equality check predicate on the "unnecessary" field. It's identical to UnnecessaryEQ.
func Unnecessary(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnnecessary), v))
	})
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
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
func StatusNotIn(vs ...Status) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// UnnecessaryEQ applies the EQ predicate on the "unnecessary" field.
func UnnecessaryEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryNEQ applies the NEQ predicate on the "unnecessary" field.
func UnnecessaryNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryIn applies the In predicate on the "unnecessary" field.
func UnnecessaryIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnnecessary), v...))
	})
}

// UnnecessaryNotIn applies the NotIn predicate on the "unnecessary" field.
func UnnecessaryNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnnecessary), v...))
	})
}

// UnnecessaryGT applies the GT predicate on the "unnecessary" field.
func UnnecessaryGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryGTE applies the GTE predicate on the "unnecessary" field.
func UnnecessaryGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryLT applies the LT predicate on the "unnecessary" field.
func UnnecessaryLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryLTE applies the LTE predicate on the "unnecessary" field.
func UnnecessaryLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryContains applies the Contains predicate on the "unnecessary" field.
func UnnecessaryContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryHasPrefix applies the HasPrefix predicate on the "unnecessary" field.
func UnnecessaryHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryHasSuffix applies the HasSuffix predicate on the "unnecessary" field.
func UnnecessaryHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryIsNil applies the IsNil predicate on the "unnecessary" field.
func UnnecessaryIsNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnnecessary)))
	})
}

// UnnecessaryNotNil applies the NotNil predicate on the "unnecessary" field.
func UnnecessaryNotNil() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnnecessary)))
	})
}

// UnnecessaryEqualFold applies the EqualFold predicate on the "unnecessary" field.
func UnnecessaryEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnnecessary), v))
	})
}

// UnnecessaryContainsFold applies the ContainsFold predicate on the "unnecessary" field.
func UnnecessaryContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnnecessary), v))
	})
}

// HasBlogPosts applies the HasEdge predicate on the "blog_posts" edge.
func HasBlogPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlogPostsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlogPostsTable, BlogPostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlogPostsWith applies the HasEdge predicate on the "blog_posts" edge with a given conditions (other predicates).
func HasBlogPostsWith(preds ...predicate.BlogPost) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlogPostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BlogPostsTable, BlogPostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProfilePic applies the HasEdge predicate on the "profile_pic" edge.
func HasProfilePic() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfilePicTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfilePicTable, ProfilePicColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfilePicWith applies the HasEdge predicate on the "profile_pic" edge with a given conditions (other predicates).
func HasProfilePicWith(preds ...predicate.Image) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfilePicInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProfilePicTable, ProfilePicColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkipEdge applies the HasEdge predicate on the "skip_edge" edge.
func HasSkipEdge() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SkipEdgeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SkipEdgeTable, SkipEdgeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkipEdgeWith applies the HasEdge predicate on the "skip_edge" edge with a given conditions (other predicates).
func HasSkipEdgeWith(preds ...predicate.SkipEdgeExample) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SkipEdgeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SkipEdgeTable, SkipEdgeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
