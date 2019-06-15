package sqlite

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/nektro/go-util/util"

	. "github.com/nektro/go-util/alias"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

type RowPragmaTableInfo struct {
	cid       int
	name      string
	typeV     string `db:"type"`
	notnull   int
	dfltValue string `db:"dflt_value"`
	pk        int
}

func Connect(path string) *DB {
	db, err := sql.Open("sqlite3", "file:"+path+"/access.db?mode=rwc&cache=shared&_busy_timeout=5000")
	util.CheckErr(err)
	db.SetMaxOpenConns(1)
	return &DB{db}
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) DB() *sql.DB {
	return db.db
}

func (db *DB) CreateTable(name string, pk []string, columns [][]string) {
	if !db.DoesTableExist(name) {
		db.Query(true, F("create table %s(%s %s)", name, pk[0], pk[1]))
		util.Log(fmt.Sprintf("Created table '%s'", name))
	}
	pti := db.QueryColumnList(name)
	for _, col := range columns {
		if !util.Contains(pti, col[0]) {
			db.Query(true, F("alter table %s add %s %s", name, col[0], col[1]))
			util.Logf("Added column '%s.%s'", name, col[0])
		}
	}
}

func (db *DB) CreateTableStruct(name string, v interface{}) {
	t := reflect.TypeOf(v)
	cols := [][]string{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		g := f.Tag.Get("sqlite")
		if len(g) > 0 {
			cols = append(cols, []string{f.Tag.Get("json"), g})
		}
	}
	db.CreateTable(name, []string{"id", "int primary key"}, cols)
}

func (db *DB) DoesTableExist(table string) bool {
	q := db.Query(false, F("select name from sqlite_master where type='table' AND name='%s';", table))
	e := q.Next()
	q.Close()
	return e
}

func (db *DB) Query(modify bool, q string) *sql.Rows {
	if modify {
		_, err := db.db.Exec(q)
		util.CheckErr(err)
		return nil
	}
	rows, err := db.db.Query(q)
	util.CheckErr(err)
	return rows
}

func (db *DB) QueryColumnList(table string) []string {
	var result []string
	rows := db.Query(false, F("pragma table_info(%s)", table))
	for rows.Next() {
		var cid int
		var name string
		var typeV string
		var notnull bool
		var dfltValue string
		var pk int
		rows.Scan(&cid, &name, &typeV, &notnull, &dfltValue, &pk)
		result = append(result, name)
	}
	rows.Close()
	return result
}

func (db *DB) QueryNextID(table string) int {
	result := -1
	rows := db.Query(false, F("select id from %s order by id desc limit 1", table))
	for rows.Next() {
		rows.Scan(&result)
	}
	rows.Close()
	return result + 1
}

func (db *DB) QueryPrepared(modify bool, q string, args ...interface{}) *sql.Rows {
	stmt, err := db.db.Prepare(q)
	util.CheckErr(err)
	if modify {
		_, err := stmt.Exec(args...)
		util.CheckErr(err)
		return nil
	}
	rows, err := stmt.Query(args...)
	util.CheckErr(err)
	return rows
}

func (db *DB) QueryDoSelectAll(table string) *sql.Rows {
	return db.Query(false, F("select * from %s", table))
}

func (db *DB) QueryDoSelect(table string, where string, search string) *sql.Rows {
	return db.QueryPrepared(false, F("select * from %s where %s = ?", table, where), search)
}

func (db *DB) QueryDoSelectAnd(table string, where string, search string, where2 string, search2 string) *sql.Rows {
	return db.QueryPrepared(false, F("select * from %s where %s = ? and %s = ?", table, where, where2), search, search2)
}

func (db *DB) QueryDoUpdate(table string, col string, value string, where string, search string) {
	db.QueryPrepared(true, F("update %s set %s = ? where %s = ?", table, col, where), value, search)
}

func (db *DB) QueryDoSelectAllOrder(table string, order string) *sql.Rows {
	return db.Query(false, F("select * from %s order by %s desc", table, order))
}
