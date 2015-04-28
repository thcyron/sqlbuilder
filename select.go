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
	group   string
	having  string
}

type sel struct {
	col  string
	dest interface{}
}

type join struct {
	sql  string
	args []interface{}
}

// From sets the table to select from.
func (s SelectStatement) From(table string) SelectStatement {
	s.table = table
	return s
}

// Map configures the statement to select column col and scan its value
// into dest. Dest may be nil if you don't want to scan the value.
func (s SelectStatement) Map(col string, dest interface{}) SelectStatement {
	if dest == nil {
		dest = nullDest
	}
	s.selects = append(s.selects, sel{col, dest})
	return s
}

// Join adds a JOIN statement to the query.
func (s SelectStatement) Join(sql string, args ...interface{}) SelectStatement {
	s.joins = append(s.joins, join{sql, args})
	return s
}

// Where adds a WHERE condition to the query. Multiple conditions are
// AND'd together.
func (s SelectStatement) Where(cond string, args ...interface{}) SelectStatement {
	s.wheres = append(s.wheres, where{cond, args})
	return s
}

// Limit sets the limit.
func (s SelectStatement) Limit(limit int) SelectStatement {
	s.limit = &limit
	return s
}

// Offset sets the offset.
func (s SelectStatement) Offset(offset int) SelectStatement {
	s.offset = &offset
	return s
}

// Order sets the ordering of the results. Only the last Order() is used
// in the query, use Order("updated_at DESC, id DESC") to order by multiple columns.
func (s SelectStatement) Order(order string) SelectStatement {
	s.order = order
	return s
}

// Group sets the GROUP BY statement. Only the last Group() is used.
func (s SelectStatement) Group(group string) SelectStatement {
	s.group = group
	return s
}

// Having sets the HAVING statement. Only the last Having() is used.
func (s SelectStatement) Having(having string) SelectStatement {
	s.having = having
	return s
}

// Lock updates the statement to lock rows using FOR UPDATE.
func (s SelectStatement) Lock() SelectStatement {
	s.lock = true
	return s
}

// Build builds the SQL query. It returns the query, the argument slice,
// and the scans slice.
func (s SelectStatement) Build() (query string, args []interface{}, scans []interface{}) {
	var cols []string
	idx := 0

	if len(s.selects) > 0 {
		for _, sel := range s.selects {
			cols = append(cols, sel.col)
			if sel.dest == nil {
				scans = append(scans, &nullDest)
			} else {
				scans = append(scans, sel.dest)
			}
		}
	} else {
		cols = append(cols, "1")
		scans = append(scans, nullDest)
	}
	query = "SELECT " + strings.Join(cols, ", ") + " FROM " + s.table

	for _, join := range s.joins {
		sql := join.sql
		for _, arg := range join.args {
			sql = strings.Replace(sql, "?", s.dbms.Placeholder(idx), 1)
			idx++
			args = append(args, arg)
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
				args = append(args, arg)
			}
			sqls = append(sqls, sql)
		}
		query += " WHERE " + strings.Join(sqls, " AND ")
	}

	if s.order != "" {
		query += " ORDER BY " + s.order
	}

	if s.group != "" {
		query += " GROUP BY " + s.group
	}

	if s.having != "" {
		query += " HAVING " + s.having
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

	return
}
