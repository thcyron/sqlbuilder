package sqlbuilder

import "strings"

type insertSet struct {
	col string
	arg interface{}
	raw bool
}

// InsertStatement represents an INSERT SQL statement.
type InsertStatement struct {
	dbms  DBMS
	table string
	sets  []insertSet
	args  []interface{}
}

// Into sets the table to insert into.
func (s *InsertStatement) Into(table string) *InsertStatement {
	s.table = table
	return s
}

// Set updates the statement to set column col to value val.
func (s *InsertStatement) Set(col string, val interface{}) *InsertStatement {
	s.sets = append(s.sets, insertSet{col, val, false})
	return s
}

// SetSQL updates the statement to set column col to the value of the raw SQL sql.
func (s *InsertStatement) SetSQL(col, sql string) *InsertStatement {
	s.sets = append(s.sets, insertSet{col, sql, true})
	return s
}

// Query builds and returns the SQL query.
func (s *InsertStatement) Query() string {
	var cols, vals []string
	idx := 0
	s.args = []interface{}{}

	for _, set := range s.sets {
		cols = append(cols, set.col)

		if set.raw {
			vals = append(vals, set.arg.(string))
		} else {
			s.args = append(s.args, set.arg)
			vals = append(vals, s.dbms.Placeholder(idx))
			idx++
		}
	}

	return "INSERT INTO " + s.table + " (" + strings.Join(cols, ", ") + ") VALUES (" + strings.Join(vals, ", ") + ")"
}

// Args returns the arguments for the placeholder in the query.
// Args() panics if Query() was not called before.
func (s *InsertStatement) Args() []interface{} {
	if s.args == nil {
		panic("sqlbuilder: must call Query() before Args()")
	}
	return s.args
}
