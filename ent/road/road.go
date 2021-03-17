// Code generated by entc, DO NOT EDIT.

package road

const (
	// Label holds the string label denoting the road type in the database.
	Label = "road"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"

	// Table holds the table name of the road in the database.
	Table = "roads"
)

// Columns holds all SQL columns for road fields.
var Columns = []string{
	FieldID,
	FieldX,
	FieldY,
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