package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestInsertMySQL(t *testing.T) {
	query, args := MySQL.Insert().
		Into("customers").
		Set("name", "John").
		Set("phone", "555").
		SetSQL("created_at", "NOW()").
		Build()

	expectedQuery := "INSERT INTO `customers` (`name`, `phone`, `created_at`) VALUES (?, ?, NOW())"
	if query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}

func TestInsertPostgres(t *testing.T) {
	query, args := Postgres.Insert().
		Into("customers").
		Set("name", "John").
		Set("phone", "555").
		SetSQL("created_at", "NOW()").
		Build()

	expectedQuery := `INSERT INTO "customers" ("name", "phone", "created_at") VALUES ($1, $2, NOW())`
	if query != expectedQuery {
		t.Errorf("bad query: %s", query)
	}

	expectedArgs := []interface{}{"John", "555"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("bad args: %v", args)
	}
}
