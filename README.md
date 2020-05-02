sqlbuilder
==========

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/thcyron/sqlbuilder/v4)
[![CI status](https://github.com/thcyron/sqlbuilder/workflows/CI/badge.svg)](https://github.com/thcyron/sqlbuilder/actions?query=workflow%3ACI)

`sqlbuilder` is a Go library for building SQL queries.

The latest stable version is [4.0.0](https://github.com/thcyron/sqlbuilder/tree/v4.0.0/).
Version 4 is identical to version 3 with added support for Go modules.

`sqlbuilder` follows [Semantic Versioning](http://semver.org/).

Usage
-----

```go
import "github.com/thcyron/sqlbuilder/v4"
```

Examples
--------

**SELECT**

```go
query, args, dest := sqlbuilder.Select().
        From("customers").
        Map("id", &customer.ID).
        Map("name", &customer.Name).
        Map("phone", &customer.Phone).
        Order("id DESC").
        Limit(1).
        Build()

err := db.QueryRow(query, args...).Scan(dest...)
```

**INSERT**

```go
query, args, dest := sqlbuilder.Insert().
        Into("customers").
        Set("name", "John").
        Set("phone", "555").
        Build()
res, err := db.Exec(query, args...)
```

**UPDATE**

```go
query, args := sqlbuilder.Update().
        Table("customers").
        Set("name", "John").
        Set("phone", "555").
        Where("id = ?", 1).
        Build()
res, err := db.Exec(query, args...)
```

**DELETE**

```go
query, args := sqlbuilder.Delete().
    From("customers").
    Where("name = ?", "John").
    Build()
res, err := db.Exec(query, args...)
```

Dialects
--------

`sqlbuilder` supports building queries for MySQL, SQLite, and Postgres databases. You
can set the default dialect with:

```go
sqlbuilder.DefaultDialect = sqlbuilder.Postgres
sqlbuilder.Select().From("...")...
```

Or you can specify the dialect explicitly:

```go
sqlbuilder.Select().Dialect(sqlbuilder.Postgres).From("...")...
```

License
-------

`sqlbuilder` is licensed under the MIT License.
