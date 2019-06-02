package sqlite

import (
	"database/sql"

	. "github.com/nektro/go-util/alias"
)

type QueryRaw struct {
	db  *DB
	sql string
}

func (db *DB) Select() *QueryRaw {
	q := QueryRaw{db, "select"}
	return &q
}

func (q *QueryRaw) Run(modify bool) *sql.Rows {
	return q.db.Query(modify, q.sql)
}

func (q *QueryRaw) Columns(cs string) *QueryRaw {
	q.sql += F(" " + cs)
	return q
}

func (q *QueryRaw) All() *QueryRaw {
	return q.Columns("*")
}

func (q *QueryRaw) From(table string) *QueryRaw {
	q.sql += F(" from %s", table)
	return q
}

func (q *QueryRaw) WhereEq(col string, search string) *QueryRaw {
	q.sql += F(" where %s = '%s'", col, search)
	return q
}

func (q *QueryRaw) AndEq(col string, search string) *QueryRaw {
	q.sql += F(" and %s = '%s'", col, search)
	return q
}

func (q *QueryRaw) OrderBy(col string) *QueryRaw {
	q.sql += F(" order by %s", col)
	return q
}

func (q *QueryRaw) Limit(from string, to string) *QueryRaw {
	q.sql += F(" limit %s,%s", from, to)
	return q
}
