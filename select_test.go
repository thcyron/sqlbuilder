package sqlbuilder

import (
	"reflect"
	"testing"
)

type customer struct {
	ID    int
	Name  string
	Phone *string
}

func TestSimpleSelect(t *testing.T) {
	c := customer{}

	q := MySQL.Select().From("customers")
	q.Map("id", &c.ID)
	q.Map("name", &c.Name)
	q.Map("phone", &c.Phone)
	q.Map("1+1 AS two", nil)

	expectedQuery := "SELECT id, name, phone, 1+1 AS two FROM customers"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedDest := []interface{}{&c.ID, &c.Name, &c.Phone, &nullDest}
	if dest := q.Dest(); !reflect.DeepEqual(dest, expectedDest) {
		t.Errorf("bad dest: %v", dest)
	}
}

func TestSimpleSelectWithLimitOffset(t *testing.T) {
	c := customer{}

	q := MySQL.Select().From("customers")
	q.Map("id", &c.ID)
	q.Map("name", &c.Name)
	q.Map("phone", &c.Phone)
	q.Limit(5)
	q.Offset(10)

	expectedQuery := "SELECT id, name, phone FROM customers LIMIT 5 OFFSET 10"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedDest := []interface{}{&c.ID, &c.Name, &c.Phone}
	if dest := q.Dest(); !reflect.DeepEqual(dest, expectedDest) {
		t.Errorf("bad dest: %v", dest)
	}
}

func TestSimpleSelectWithJoins(t *testing.T) {
	c := customer{}

	q := MySQL.Select().From("customers")
	q.Map("id", &c.ID)
	q.Map("name", &c.Name)
	q.Map("phone", &c.Phone)
	q.Join("INNER JOIN orders ON orders.customer_id = customers.id")
	q.Join("LEFT JOIN items ON items.order_id = orders.id")

	expectedQuery := "SELECT id, name, phone FROM customers INNER JOIN orders ON orders.customer_id = customers.id LEFT JOIN items ON items.order_id = orders.id"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}
}

func TestSelectWithWhereMySQL(t *testing.T) {
	c := customer{}

	q := MySQL.Select().From("customers")
	q.Map("id", &c.ID)
	q.Map("name", &c.Name)
	q.Map("phone", &c.Phone)
	q.Where("id = ? AND name IS NOT NULL", 9)

	expectedQuery := "SELECT id, name, phone FROM customers WHERE id = ? AND name IS NOT NULL"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{9}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestSelectWithGroupMySQL(t *testing.T) {
	var count uint
	q := MySQL.Select().From("customers").Map("COUNT(*)", &count).Group("city")
	expectedQuery := "SELECT COUNT(*) FROM customers GROUP BY city"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}
}

func TestSelectWithWherePostgres(t *testing.T) {
	c := customer{}

	q := Postgres.Select().From("customers")
	q.Map("id", &c.ID)
	q.Map("name", &c.Name)
	q.Map("phone", &c.Phone)
	q.Where("id = ? AND name IS NOT NULL", 9)

	expectedQuery := "SELECT id, name, phone FROM customers WHERE id = $1 AND name IS NOT NULL"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{9}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}
