// Code generated by ent, DO NOT EDIT.

package ent

import (
	"abc/internal/data/ent/words"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Words is the model entity for the Words schema.
type Words struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Words) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case words.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Words fields.
func (w *Words) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case words.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Words.
// This includes values selected through modifiers, order, etc.
func (w *Words) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Words.
// Note that you need to call Words.Unwrap() before calling this method if this Words
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Words) Update() *WordsUpdateOne {
	return NewWordsClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Words entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Words) Unwrap() *Words {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Words is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Words) String() string {
	var builder strings.Builder
	builder.WriteString("Words(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WordsSlice is a parsable slice of Words.
type WordsSlice []*Words
