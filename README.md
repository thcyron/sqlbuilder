sqlbuilder
==========

`sqlbuilder` is a Go library for building SQL queries.

Examples
--------

**SELECT**

```go
b := sqlbuilder.Select().From("customers").
        Select("id", &customer.ID).
        Select("name", &customer.Name).
        Select("phone", &customer.Phone).
        Order("id DESC").
        Limit(1)
err := db.QueryRow(b.Query(), b.Args()...).Scan(b.Dest()...)
```

**INSERT**

```go
b := sqlbuilder.Insert("customers").
        Set("name", "John").
        Set("phone", "555")
err := db.Exec(b.Query(), b.Args()...)
```

**UPDATE**

```go
b := sqlbuilder.Update("customers").
        Set("name", "John").
        Set("phone", "555").
        Where("id = ?", 1)
err := db.Exec(b.Query(), b.Args()...)
```

Supported DBMS
--------------

`sqlbuilder` supports building queries for MySQL and Postgres databases. You
can set the default DBMS used by the package-level Select, Update and Insert
functions with:

```go
sqlbuilder.DefaultDBMS = Postgres
sqlbuilder.Select("...")
```

or you can specify the DBMS explicitly:

```go
sqlbuilder.Postgres.Select("...")
```

Documentation
-------------

Documentation is available at [Godoc](http://godoc.org/github.com/thcyron/sqlbuilder).

License
-------

`sqlbuilder` is licensed under the MIT license.
