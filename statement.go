package sqlbuilder

import "database/sql"

// A Statement has a query and arguments.
type Statement interface {
	Query() string
	Args() []interface{}
}

// ReturnStatement is a Statement that additionally returns values.
type ReturnStatement interface {
	Statement
	Dest() []interface{}
}

// ExecDB runs Exec() with the given datebase and statement.
func ExecDB(db *sql.DB, s Statement) (sql.Result, error) {
	return db.Exec(s.Query(), s.Args()...)
}

// QueryDB runs Query() with the given datebase and statement.
func QueryDB(db *sql.DB, s ReturnStatement) (*sql.Rows, error) {
	return db.Query(s.Query(), s.Args()...)
}

// QueryRowDB runs QueryRow() and Scan() with the given datebase and statement.
func QueryRowDB(db *sql.DB, s ReturnStatement) error {
	return db.QueryRow(s.Query(), s.Args()...).Scan(s.Dest()...)
}

// ExecTx runs Exec() with the given transaction and statement.
func ExecTx(tx *sql.Tx, s Statement) (sql.Result, error) {
	return tx.Exec(s.Query(), s.Args()...)
}

// QueryTx runs Query() with the given transaction and statement.
func QueryTx(tx *sql.Tx, s ReturnStatement) (*sql.Rows, error) {
	return tx.Query(s.Query(), s.Args()...)
}

// QueryRowTx runs QueryRow() and Scan() with the given transaction
// and statement.
func QueryRowTx(tx *sql.Tx, s ReturnStatement) error {
	return tx.QueryRow(s.Query(), s.Args()...).Scan(s.Dest()...)
}
