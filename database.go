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
		panic("sqlbuilder: unknown DBMS")
	}
}

// Select returns a SELECT statement.
func (dbms DBMS) Select() SelectStatement {
	return SelectStatement{dbms: dbms}
}

// Insert returns an INSERT statement.
func (dbms DBMS) Insert() InsertStatement {
	return InsertStatement{dbms: dbms}
}

// Update returns an UPDATE statement.
func (dbms DBMS) Update() UpdateStatement {
	return UpdateStatement{dbms: dbms}
}

// DefaultDBMS is the DBMS used by the package-level Select, Insert and Update functions.
var DefaultDBMS = MySQL

// Select returns a SELECT statement using the default Database.
func Select() SelectStatement {
	return DefaultDBMS.Select()
}

// Insert returns an INSERT statement using the default Database.
func Insert() InsertStatement {
	return DefaultDBMS.Insert()
}

// Update returns an UPDATE statement using the default Database.
func Update() UpdateStatement {
	return DefaultDBMS.Update()
}
