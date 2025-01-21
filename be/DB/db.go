package DB

import "github.com/go-pg/pg"

func DBinit() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:56363",
		User:     "postgres",
		Password: "114514",
		Database: "testdb",
	})
	return db
}

var Db *pg.DB
