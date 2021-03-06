// Code generated by entc, DO NOT EDIT.

package settlement

const (
	// Label holds the string label denoting the settlement type in the database.
	Label = "settlement"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"
	// FieldIsCity holds the string denoting the is_city field in the database.
	FieldIsCity = "is_city"

	// EdgeHexes holds the string denoting the hexes edge name in mutations.
	EdgeHexes = "hexes"
	// EdgeHarbor holds the string denoting the harbor edge name in mutations.
	EdgeHarbor = "harbor"

	// Table holds the table name of the settlement in the database.
	Table = "settlements"
	// HexesTable is the table the holds the hexes relation/edge. The primary key declared below.
	HexesTable = "settlement_hexes"
	// HexesInverseTable is the table name for the Hex entity.
	// It exists in this package in order to avoid circular dependency with the "hex" package.
	HexesInverseTable = "hexes"
	// HarborTable is the table the holds the harbor relation/edge.
	HarborTable = "harbors"
	// HarborInverseTable is the table name for the Harbor entity.
	// It exists in this package in order to avoid circular dependency with the "harbor" package.
	HarborInverseTable = "harbors"
	// HarborColumn is the table column denoting the harbor relation/edge.
	HarborColumn = "settlement_harbor"
)

// Columns holds all SQL columns for settlement fields.
var Columns = []string{
	FieldID,
	FieldX,
	FieldY,
	FieldIsCity,
}

var (
	// HexesPrimaryKey and HexesColumn2 are the table columns denoting the
	// primary key for the hexes relation (M2M).
	HexesPrimaryKey = []string{"settlement_id", "hex_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
