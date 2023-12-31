package dat

import "github.com/unix-world/go-modules/mgutz/dat/common"

// Dialect is the active SQLDialect.
var Dialect SQLDialect

// SQLDialect represents a vendor specific SQL dialect.
type SQLDialect interface {
	// WriteStringLiteral writes a string literal.
	WriteStringLiteral(buf common.BufferWriter, value string)
	// WriteIdentifier writes a quoted identifer such as a column or table.
	WriteIdentifier(buf common.BufferWriter, column string)
}
