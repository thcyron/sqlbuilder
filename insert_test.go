package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestInsertMySQL(t *testing.T) {
	q := MySQL.Insert().Into("customers")
	q.Set("name", "John")
	q.Set("phone", "555")
	q.SetSQL("created_at", "NOW()")

	expectedQuery := "INSERT INTO customers (name, phone, created_at) VALUES (?, ?, NOW())"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestInsertPostgres(t *testing.T) {
	q := Postgres.Insert().Into("customers")
	q.Set("name", "John")
	q.Set("phone", "555")
	q.SetSQL("created_at", "NOW()")

	expectedQuery := "INSERT INTO customers (name, phone, created_at) VALUES ($1, $2, NOW())"
	if query := q.Query(); query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if args := q.Args(); !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}
