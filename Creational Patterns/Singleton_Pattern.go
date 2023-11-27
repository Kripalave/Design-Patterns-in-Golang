package singleton

import (
 "database/sql"
 "sync"

 _ "github.com/go-sql-driver/mysql"
)

type Database struct {
 conn *sql.DB
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
 once.Do(func() {
  db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/mydatabase")
  if err != nil {
   panic(err.Error())
  }
  instance = &Database{conn: db}
 })
 return instance
}

func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
 return d.conn.Query(query, args...)
}

func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
 return d.conn.Exec(query, args...)
}
