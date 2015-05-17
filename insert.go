package sqlbuilder

import "strings"

type insertSet struct {
	col string
	arg interface{}
	raw bool
}

// InsertStatement represents an INSERT statement.
type InsertStatement struct {
	dbms  DBMS
	table string
	sets  []insertSet
}

// Into returns a new statement with the table to insert into set to 'table'.
func (s InsertStatement) Into(table string) InsertStatement {
	s.table = table
	return s
}

// Set returns a new statement with column 'col' set to value 'val'.
func (s InsertStatement) Set(col string, val interface{}) InsertStatement {
	s.sets = append(s.sets, insertSet{col, val, false})
	return s
}

// SetSQL returns a new statement with column 'col' set to the raw SQL expression 'sql'.
func (s InsertStatement) SetSQL(col, sql string) InsertStatement {
	s.sets = append(s.sets, insertSet{col, sql, true})
	return s
}

// Build builds the SQL query. It returns the SQL query and the argument slice.
func (s InsertStatement) Build() (query string, args []interface{}) {
	var cols, vals []string
	idx := 0

	for _, set := range s.sets {
		cols = append(cols, set.col)

		if set.raw {
			vals = append(vals, set.arg.(string))
		} else {
			args = append(args, set.arg)
			vals = append(vals, s.dbms.Placeholder(idx))
			idx++
		}
	}

	query = "INSERT INTO " + s.table + " (" + strings.Join(cols, ", ") + ") VALUES (" + strings.Join(vals, ", ") + ")"
	return
}
