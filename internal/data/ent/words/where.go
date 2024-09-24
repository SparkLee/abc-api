// Code generated by ent, DO NOT EDIT.

package words

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sparklee/abc-api/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Words {
	return predicate.Words(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Words {
	return predicate.Words(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Words {
	return predicate.Words(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Words {
	return predicate.Words(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Words {
	return predicate.Words(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Words {
	return predicate.Words(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Words {
	return predicate.Words(sql.FieldLTE(FieldID, id))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldGroup, v))
}

// Word applies equality check predicate on the "word" field. It's identical to WordEQ.
func Word(v string) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldWord, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldUpdatedAt, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.Words {
	return predicate.Words(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.Words {
	return predicate.Words(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.Words {
	return predicate.Words(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.Words {
	return predicate.Words(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.Words {
	return predicate.Words(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.Words {
	return predicate.Words(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.Words {
	return predicate.Words(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.Words {
	return predicate.Words(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.Words {
	return predicate.Words(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.Words {
	return predicate.Words(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.Words {
	return predicate.Words(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.Words {
	return predicate.Words(sql.FieldContainsFold(FieldGroup, v))
}

// WordEQ applies the EQ predicate on the "word" field.
func WordEQ(v string) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldWord, v))
}

// WordNEQ applies the NEQ predicate on the "word" field.
func WordNEQ(v string) predicate.Words {
	return predicate.Words(sql.FieldNEQ(FieldWord, v))
}

// WordIn applies the In predicate on the "word" field.
func WordIn(vs ...string) predicate.Words {
	return predicate.Words(sql.FieldIn(FieldWord, vs...))
}

// WordNotIn applies the NotIn predicate on the "word" field.
func WordNotIn(vs ...string) predicate.Words {
	return predicate.Words(sql.FieldNotIn(FieldWord, vs...))
}

// WordGT applies the GT predicate on the "word" field.
func WordGT(v string) predicate.Words {
	return predicate.Words(sql.FieldGT(FieldWord, v))
}

// WordGTE applies the GTE predicate on the "word" field.
func WordGTE(v string) predicate.Words {
	return predicate.Words(sql.FieldGTE(FieldWord, v))
}

// WordLT applies the LT predicate on the "word" field.
func WordLT(v string) predicate.Words {
	return predicate.Words(sql.FieldLT(FieldWord, v))
}

// WordLTE applies the LTE predicate on the "word" field.
func WordLTE(v string) predicate.Words {
	return predicate.Words(sql.FieldLTE(FieldWord, v))
}

// WordContains applies the Contains predicate on the "word" field.
func WordContains(v string) predicate.Words {
	return predicate.Words(sql.FieldContains(FieldWord, v))
}

// WordHasPrefix applies the HasPrefix predicate on the "word" field.
func WordHasPrefix(v string) predicate.Words {
	return predicate.Words(sql.FieldHasPrefix(FieldWord, v))
}

// WordHasSuffix applies the HasSuffix predicate on the "word" field.
func WordHasSuffix(v string) predicate.Words {
	return predicate.Words(sql.FieldHasSuffix(FieldWord, v))
}

// WordEqualFold applies the EqualFold predicate on the "word" field.
func WordEqualFold(v string) predicate.Words {
	return predicate.Words(sql.FieldEqualFold(FieldWord, v))
}

// WordContainsFold applies the ContainsFold predicate on the "word" field.
func WordContainsFold(v string) predicate.Words {
	return predicate.Words(sql.FieldContainsFold(FieldWord, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Words {
	return predicate.Words(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Words {
	return predicate.Words(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Words {
	return predicate.Words(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Words {
	return predicate.Words(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Words {
	return predicate.Words(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Words) predicate.Words {
	return predicate.Words(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Words) predicate.Words {
	return predicate.Words(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Words) predicate.Words {
	return predicate.Words(sql.NotPredicates(p))
}