package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	query, args := Delete().
		Dialect(MySQL).
		From("customers").
		Build()

	expectedQuery := "DELETE FROM customers"
	if query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	if args != nil {
		t.Errorf("bad args: %v", args)
	}
}

func TestDeleteWhere(t *testing.T) {
	query, args := Delete().
		Dialect(MySQL).
		From("customers").
		Where("name = ?", "John").
		Where("phone = ?", "555").
		Build()

	expectedQuery := "DELETE FROM customers WHERE (name = ?) AND (phone = ?)"
	if query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}
