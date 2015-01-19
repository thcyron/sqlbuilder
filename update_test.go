package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestUpdateMySQL(t *testing.T) {
	q := MySQL.Update().Table("customers")
	q.Set("name", "John")
	q.Set("phone", "555")

	expectedQuery := "UPDATE customers SET name = ?, phone = ?"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestUpdatePostgres(t *testing.T) {
	q := Postgres.Update().Table("customers")
	q.Set("name", "John")
	q.Set("phone", "555")

	expectedQuery := "UPDATE customers SET name = $1, phone = $2"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestUpdateWithWhereMySQL(t *testing.T) {
	q := MySQL.Update().Table("customers")
	q.Set("name", "John")
	q.Set("phone", "555")
	q.Where("id = ?", 9)

	expectedQuery := "UPDATE customers SET name = ?, phone = ? WHERE id = ?"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555", 9}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestUpdateWithWherePostgres(t *testing.T) {
	q := Postgres.Update().Table("customers")
	q.Set("name", "John")
	q.Set("phone", "555")
	q.Where("id = ?", 9)

	expectedQuery := "UPDATE customers SET name = $1, phone = $2 WHERE id = $3"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555", 9}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}
