// Code generated by ent, DO NOT EDIT.

package api

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.API {
	return predicate.API(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.API {
	return predicate.API(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.API {
	return predicate.API(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.API {
	return predicate.API(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.API {
	return predicate.API(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.API {
	return predicate.API(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.API {
	return predicate.API(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.API {
	return predicate.API(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.API {
	return predicate.API(sql.FieldEQ(FieldUpdatedAt, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldPath, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldDescription, v))
}

// APIGroup applies equality check predicate on the "api_group" field. It's identical to APIGroupEQ.
func APIGroup(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldAPIGroup, v))
}

// ServiceName applies equality check predicate on the "service_name" field. It's identical to ServiceNameEQ.
func ServiceName(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldServiceName, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldMethod, v))
}

// IsRequired applies equality check predicate on the "is_required" field. It's identical to IsRequiredEQ.
func IsRequired(v bool) predicate.API {
	return predicate.API(sql.FieldEQ(FieldIsRequired, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.API {
	return predicate.API(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.API {
	return predicate.API(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.API {
	return predicate.API(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.API {
	return predicate.API(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.API {
	return predicate.API(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.API {
	return predicate.API(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.API {
	return predicate.API(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.API {
	return predicate.API(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.API {
	return predicate.API(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.API {
	return predicate.API(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.API {
	return predicate.API(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.API {
	return predicate.API(sql.FieldLTE(FieldUpdatedAt, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.API {
	return predicate.API(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.API {
	return predicate.API(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.API {
	return predicate.API(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.API {
	return predicate.API(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.API {
	return predicate.API(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.API {
	return predicate.API(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.API {
	return predicate.API(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.API {
	return predicate.API(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.API {
	return predicate.API(sql.FieldContainsFold(FieldPath, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.API {
	return predicate.API(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.API {
	return predicate.API(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.API {
	return predicate.API(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.API {
	return predicate.API(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.API {
	return predicate.API(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.API {
	return predicate.API(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.API {
	return predicate.API(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.API {
	return predicate.API(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.API {
	return predicate.API(sql.FieldContainsFold(FieldDescription, v))
}

// APIGroupEQ applies the EQ predicate on the "api_group" field.
func APIGroupEQ(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldAPIGroup, v))
}

// APIGroupNEQ applies the NEQ predicate on the "api_group" field.
func APIGroupNEQ(v string) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldAPIGroup, v))
}

// APIGroupIn applies the In predicate on the "api_group" field.
func APIGroupIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldIn(FieldAPIGroup, vs...))
}

// APIGroupNotIn applies the NotIn predicate on the "api_group" field.
func APIGroupNotIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldAPIGroup, vs...))
}

// APIGroupGT applies the GT predicate on the "api_group" field.
func APIGroupGT(v string) predicate.API {
	return predicate.API(sql.FieldGT(FieldAPIGroup, v))
}

// APIGroupGTE applies the GTE predicate on the "api_group" field.
func APIGroupGTE(v string) predicate.API {
	return predicate.API(sql.FieldGTE(FieldAPIGroup, v))
}

// APIGroupLT applies the LT predicate on the "api_group" field.
func APIGroupLT(v string) predicate.API {
	return predicate.API(sql.FieldLT(FieldAPIGroup, v))
}

// APIGroupLTE applies the LTE predicate on the "api_group" field.
func APIGroupLTE(v string) predicate.API {
	return predicate.API(sql.FieldLTE(FieldAPIGroup, v))
}

// APIGroupContains applies the Contains predicate on the "api_group" field.
func APIGroupContains(v string) predicate.API {
	return predicate.API(sql.FieldContains(FieldAPIGroup, v))
}

// APIGroupHasPrefix applies the HasPrefix predicate on the "api_group" field.
func APIGroupHasPrefix(v string) predicate.API {
	return predicate.API(sql.FieldHasPrefix(FieldAPIGroup, v))
}

// APIGroupHasSuffix applies the HasSuffix predicate on the "api_group" field.
func APIGroupHasSuffix(v string) predicate.API {
	return predicate.API(sql.FieldHasSuffix(FieldAPIGroup, v))
}

// APIGroupEqualFold applies the EqualFold predicate on the "api_group" field.
func APIGroupEqualFold(v string) predicate.API {
	return predicate.API(sql.FieldEqualFold(FieldAPIGroup, v))
}

// APIGroupContainsFold applies the ContainsFold predicate on the "api_group" field.
func APIGroupContainsFold(v string) predicate.API {
	return predicate.API(sql.FieldContainsFold(FieldAPIGroup, v))
}

// ServiceNameEQ applies the EQ predicate on the "service_name" field.
func ServiceNameEQ(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldServiceName, v))
}

// ServiceNameNEQ applies the NEQ predicate on the "service_name" field.
func ServiceNameNEQ(v string) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldServiceName, v))
}

// ServiceNameIn applies the In predicate on the "service_name" field.
func ServiceNameIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldIn(FieldServiceName, vs...))
}

// ServiceNameNotIn applies the NotIn predicate on the "service_name" field.
func ServiceNameNotIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldServiceName, vs...))
}

// ServiceNameGT applies the GT predicate on the "service_name" field.
func ServiceNameGT(v string) predicate.API {
	return predicate.API(sql.FieldGT(FieldServiceName, v))
}

// ServiceNameGTE applies the GTE predicate on the "service_name" field.
func ServiceNameGTE(v string) predicate.API {
	return predicate.API(sql.FieldGTE(FieldServiceName, v))
}

// ServiceNameLT applies the LT predicate on the "service_name" field.
func ServiceNameLT(v string) predicate.API {
	return predicate.API(sql.FieldLT(FieldServiceName, v))
}

// ServiceNameLTE applies the LTE predicate on the "service_name" field.
func ServiceNameLTE(v string) predicate.API {
	return predicate.API(sql.FieldLTE(FieldServiceName, v))
}

// ServiceNameContains applies the Contains predicate on the "service_name" field.
func ServiceNameContains(v string) predicate.API {
	return predicate.API(sql.FieldContains(FieldServiceName, v))
}

// ServiceNameHasPrefix applies the HasPrefix predicate on the "service_name" field.
func ServiceNameHasPrefix(v string) predicate.API {
	return predicate.API(sql.FieldHasPrefix(FieldServiceName, v))
}

// ServiceNameHasSuffix applies the HasSuffix predicate on the "service_name" field.
func ServiceNameHasSuffix(v string) predicate.API {
	return predicate.API(sql.FieldHasSuffix(FieldServiceName, v))
}

// ServiceNameEqualFold applies the EqualFold predicate on the "service_name" field.
func ServiceNameEqualFold(v string) predicate.API {
	return predicate.API(sql.FieldEqualFold(FieldServiceName, v))
}

// ServiceNameContainsFold applies the ContainsFold predicate on the "service_name" field.
func ServiceNameContainsFold(v string) predicate.API {
	return predicate.API(sql.FieldContainsFold(FieldServiceName, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.API {
	return predicate.API(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.API {
	return predicate.API(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.API {
	return predicate.API(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.API {
	return predicate.API(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.API {
	return predicate.API(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.API {
	return predicate.API(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.API {
	return predicate.API(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.API {
	return predicate.API(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.API {
	return predicate.API(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.API {
	return predicate.API(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.API {
	return predicate.API(sql.FieldContainsFold(FieldMethod, v))
}

// IsRequiredEQ applies the EQ predicate on the "is_required" field.
func IsRequiredEQ(v bool) predicate.API {
	return predicate.API(sql.FieldEQ(FieldIsRequired, v))
}

// IsRequiredNEQ applies the NEQ predicate on the "is_required" field.
func IsRequiredNEQ(v bool) predicate.API {
	return predicate.API(sql.FieldNEQ(FieldIsRequired, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.API) predicate.API {
	return predicate.API(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.API) predicate.API {
	return predicate.API(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.API) predicate.API {
	return predicate.API(sql.NotPredicates(p))
}
