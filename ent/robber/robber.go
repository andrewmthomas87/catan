// Code generated by entc, DO NOT EDIT.

package robber

const (
	// Label holds the string label denoting the robber type in the database.
	Label = "robber"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// EdgeHex holds the string denoting the hex edge name in mutations.
	EdgeHex = "hex"

	// Table holds the table name of the robber in the database.
	Table = "robbers"
	// HexTable is the table the holds the hex relation/edge.
	HexTable = "hexes"
	// HexInverseTable is the table name for the Hex entity.
	// It exists in this package in order to avoid circular dependency with the "hex" package.
	HexInverseTable = "hexes"
	// HexColumn is the table column denoting the hex relation/edge.
	HexColumn = "robber_hex"
)

// Columns holds all SQL columns for robber fields.
var Columns = []string{
	FieldID,
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