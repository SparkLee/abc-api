// Code generated by ent, DO NOT EDIT.

package words

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the words type in the database.
	Label = "words"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldWord holds the string denoting the word field in the database.
	FieldWord = "word"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the words in the database.
	Table = "words"
)

// Columns holds all SQL columns for words fields.
var Columns = []string{
	FieldID,
	FieldGroup,
	FieldWord,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)

// OrderOption defines the ordering options for the Words queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGroup orders the results by the group field.
func ByGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroup, opts...).ToFunc()
}

// ByWord orders the results by the word field.
func ByWord(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWord, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
