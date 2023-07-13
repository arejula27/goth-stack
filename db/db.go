package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/arejula27/bs/util"
	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var Bun *bun.DB

func Init() {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		name     = os.Getenv("DB_NAME")
		uri      = fmt.Sprintf("postgresql://%s:%s@%s/%s", user, password, host, name)
		sqldb    = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
		once     = sync.Once{}
	)
	fmt.Println(uri)
	once.Do(func() {
		Bun = bun.NewDB(sqldb, pgdialect.New())
		if util.IsDev() {
			Bun.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
		}
	})
}
