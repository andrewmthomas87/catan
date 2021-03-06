// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/andrewmthomas87/catan/ent/hex"
	"github.com/andrewmthomas87/catan/ent/robber"
)

// Robber is the model entity for the Robber schema.
type Robber struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RobberQuery when eager-loading is set.
	Edges RobberEdges `json:"edges"`
}

// RobberEdges holds the relations/edges for other nodes in the graph.
type RobberEdges struct {
	// Hex holds the value of the hex edge.
	Hex *Hex `json:"hex,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HexOrErr returns the Hex value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RobberEdges) HexOrErr() (*Hex, error) {
	if e.loadedTypes[0] {
		if e.Hex == nil {
			// The edge hex was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: hex.Label}
		}
		return e.Hex, nil
	}
	return nil, &NotLoadedError{edge: "hex"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Robber) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case robber.FieldID:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Robber", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Robber fields.
func (r *Robber) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case robber.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		}
	}
	return nil
}

// QueryHex queries the "hex" edge of the Robber entity.
func (r *Robber) QueryHex() *HexQuery {
	return (&RobberClient{config: r.config}).QueryHex(r)
}

// Update returns a builder for updating this Robber.
// Note that you need to call Robber.Unwrap() before calling this method if this Robber
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Robber) Update() *RobberUpdateOne {
	return (&RobberClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Robber entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Robber) Unwrap() *Robber {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Robber is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Robber) String() string {
	var builder strings.Builder
	builder.WriteString("Robber(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Robbers is a parsable slice of Robber.
type Robbers []*Robber

func (r Robbers) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
