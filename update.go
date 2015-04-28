package sqlbuilder

import (
	"strings"
)

type updateSet struct {
	col string
	arg interface{}
	raw bool
}

// UpdateStatement represents an UPDATE statement.
type UpdateStatement struct {
	dbms   DBMS
	table  string
	sets   []updateSet
	wheres []where
	args   []interface{}
}

// Table sets the table to update.
func (s UpdateStatement) Table(table string) UpdateStatement {
	s.table = table
	return s
}

// Set updates the query to set column col to arg.
func (s UpdateStatement) Set(col string, arg interface{}) UpdateStatement {
	s.sets = append(s.sets, updateSet{col: col, arg: arg, raw: false})
	return s
}

// SetSQL updates the query to set column col to the value of the SQL expression sql.
func (s UpdateStatement) SetSQL(col string, sql string) UpdateStatement {
	s.sets = append(s.sets, updateSet{col: col, arg: sql, raw: true})
	return s
}

// Where adds a where condition to the query. Multiple conditions are AND’d together.
func (s UpdateStatement) Where(cond string, args ...interface{}) UpdateStatement {
	s.wheres = append(s.wheres, where{cond, args})
	return s
}

// Build builds the SQL query. It returns the query and the argument slice.
func (s UpdateStatement) Build() (query string, args []interface{}) {
	if len(s.sets) == 0 {
		panic("sqlbuilder: no columns set")
	}

	query = "UPDATE " + s.table + " SET "
	var sets []string
	idx := 0

	for _, set := range s.sets {
		var arg string
		if set.raw {
			arg = set.arg.(string)
		} else {
			arg = s.dbms.Placeholder(idx)
			idx++
			args = append(args, set.arg)
		}
		sets = append(sets, set.col+" = "+arg)
	}
	query += strings.Join(sets, ", ")

	if len(s.wheres) > 0 {
		var sqls []string

		for _, w := range s.wheres {
			sql := w.sql
			for _, arg := range w.args {
				p := s.dbms.Placeholder(idx)
				idx++
				sql = strings.Replace(sql, "?", p, 1)
				sqls = append(sqls, sql)
				args = append(args, arg)
			}
		}

		query += " WHERE " + strings.Join(sqls, " AND ")
	}

	return
}
