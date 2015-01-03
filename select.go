package sqlbuilder

import (
	"strconv"
	"strings"
)

var nullDest interface{}

// SelectStatement represents a SELECT statement.
type SelectStatement struct {
	dbms    DBMS
	table   string
	selects []sel
	joins   []join
	wheres  []where
	lock    bool
	limit   *int
	offset  *int
	order   string
	args    []interface{}
	dests   []interface{}
}

type sel struct {
	col  string
	dest interface{}
}

type join struct {
	sql  string
	args []interface{}
}

// Select updates the query to select column col and scan its value into dest.
// Dest may be nil if you don’t care about the value.
func (s *SelectStatement) Select(col string, dest interface{}) *SelectStatement {
	if dest == nil {
		dest = nullDest
	}
	s.selects = append(s.selects, sel{col, dest})
	return s
}

// From sets the table to select from.
func (s *SelectStatement) From(table string) *SelectStatement {
	s.table = table
	return s
}

// Join adds a JOIN statement to the query. The statement must be complete JOIN
// statement like ‘INNER JOIN foo ON foo.id = bar.foo_id’.
func (s *SelectStatement) Join(sql string, args ...interface{}) *SelectStatement {
	s.joins = append(s.joins, join{sql, args})
	return s
}

// Where adds a where condition to the query. Multiple conditions are AND’d together.
func (s *SelectStatement) Where(cond string, args ...interface{}) *SelectStatement {
	s.wheres = append(s.wheres, where{cond, args})
	return s
}

// Limit sets the limit.
func (s *SelectStatement) Limit(limit int) *SelectStatement {
	s.limit = &limit
	return s
}

// Offset sets the offset.
func (s *SelectStatement) Offset(offset int) *SelectStatement {
	s.offset = &offset
	return s
}

// Order sets the ordering of the results. Only the last Order() is used
// in the query, use Order("updated_at DESC, id DESC") to order by multiple columns.
func (s *SelectStatement) Order(order string) *SelectStatement {
	s.order = order
	return s
}

// Lock updates the statement to lock rows using FOR UPDATE.
func (s *SelectStatement) Lock() *SelectStatement {
	s.lock = true
	return s
}

// Query builds and returns the SQL query.
func (s *SelectStatement) Query() string {
	var cols []string
	idx := 0
	s.args = []interface{}{}
	s.dests = []interface{}{}

	if len(s.selects) > 0 {
		for _, sel := range s.selects {
			cols = append(cols, sel.col)
			if sel.dest == nil {
				s.dests = append(s.dests, &nullDest)
			} else {
				s.dests = append(s.dests, sel.dest)
			}
		}
	} else {
		cols = append(cols, "1")
		s.dests = append(s.dests, nullDest)
	}
	query := "SELECT " + strings.Join(cols, ", ") + " FROM " + s.table

	for _, join := range s.joins {
		sql := join.sql
		for _, arg := range join.args {
			sql = strings.Replace(sql, "?", s.dbms.Placeholder(idx), 1)
			idx++
			s.args = append(s.args, arg)
		}
		query += " " + sql
	}

	if len(s.wheres) > 0 {
		var sqls []string
		for _, where := range s.wheres {
			sql := where.sql
			for _, arg := range where.args {
				sql = strings.Replace(sql, "?", s.dbms.Placeholder(idx), 1)
				idx++
				s.args = append(s.args, arg)
			}
			sqls = append(sqls, sql)
		}
		query += " WHERE " + strings.Join(sqls, " AND ")
	}

	if s.order != "" {
		query += " ORDER BY " + s.order
	}

	if s.limit != nil {
		query += " LIMIT " + strconv.Itoa(*s.limit)
	}

	if s.offset != nil {
		query += " OFFSET " + strconv.Itoa(*s.offset)
	}

	if s.lock {
		query += " FOR UPDATE"
	}

	return query
}

// Args returns the arguments for the placeholder in the query.
// Args() panics if Query() was not called before.
func (s *SelectStatement) Args() []interface{} {
	if s.args == nil {
		panic("sqlbuilder: must call Query() before Args()")
	}
	return s.args
}

// Dest returns the destinations to scan selected columns to.
// Dest() panics if Query() was not called before.
func (s *SelectStatement) Dest() []interface{} {
	if s.dests == nil {
		panic("sqlbuilder: must call Query() before Dest()")
	}
	return s.dests
}
