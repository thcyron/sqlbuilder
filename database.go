package sqlbuilder

import "strconv"

// DBMS represents a DBMS.
type DBMS int

const (
	MySQL    DBMS = iota // MySQL
	Postgres             // Postgres
)

// Placeholder returns the placeholder string for the given index.
func (dbms DBMS) Placeholder(idx int) string {
	switch dbms {
	case MySQL:
		return "?"
	case Postgres:
		return "$" + strconv.Itoa(idx+1)
	default:
		panic("unknown DBMS")
	}
}

// Select returns a SELECT statement.
func (dbms DBMS) Select(table string) *SelectStatement {
	return &SelectStatement{
		dbms:  dbms,
		table: table,
	}
}

// Insert returns an INSERT statement.
func (dbms DBMS) Insert(table string) *InsertStatement {
	return &InsertStatement{
		dbms:  dbms,
		table: table,
	}
}

// Update returns an UPDATE statement.
func (dbms DBMS) Update(table string) *UpdateStatement {
	return &UpdateStatement{
		dbms:  dbms,
		table: table,
	}
}

// DefaultDBMS is the DBMS used by the package-level Select, Insert and Update functions.
var DefaultDatabase = MySQL

// Select returns a SELECT statement using the default Database.
func Select(table string) *SelectStatement {
	return DefaultDatabase.Select(table)
}

// Insert returns an INSERT statement using the default Database.
func Insert(table string) *InsertStatement {
	return DefaultDatabase.Insert(table)
}

// Update returns an UPDATE statement using the default Database.
func Update(table string) *UpdateStatement {
	return DefaultDatabase.Update(table)
}
