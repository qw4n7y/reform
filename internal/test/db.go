// Package test provides shared testing utilities.
package test

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	sqlite3Driver "github.com/mattn/go-sqlite3"

	"github.com/qw4n7y/reform"
	"github.com/qw4n7y/reform/dialects"
	"github.com/qw4n7y/reform/dialects/mssql" //nolint:staticcheck
	"github.com/qw4n7y/reform/dialects/mysql"
	"github.com/qw4n7y/reform/dialects/postgresql"
	"github.com/qw4n7y/reform/dialects/sqlite3"
	"github.com/qw4n7y/reform/dialects/sqlserver"
)

//nolint:gochecknoglobals
var (
	sqlite3RegisterOnce sync.Once
	inspectOnce         sync.Once
)

// ConnectToTestDB returns open and prepared connection to test DB.
func ConnectToTestDB() *reform.DB {
	driver := strings.TrimSpace(os.Getenv("REFORM_TEST_DRIVER"))
	source := strings.TrimSpace(os.Getenv("REFORM_TEST_SOURCE"))
	if driver == "" || source == "" {
		log.Fatal("no driver or source, set REFORM_TEST_DRIVER and REFORM_TEST_SOURCE")
	}

	// register custom function "sleep" for context tests
	if driver == "sqlite3" {
		driver = "sqlite3_with_sleep"

		sqlite3RegisterOnce.Do(func() {
			sleep := func(nsec int64) (int64, error) {
				time.Sleep(time.Duration(nsec))
				return nsec, nil
			}
			sql.Register(driver, &sqlite3Driver.SQLiteDriver{
				ConnectHook: func(conn *sqlite3Driver.SQLiteConn) error {
					return conn.RegisterFunc("sleep", sleep, false)
				},
			})
		})
	}

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}

	// Use single connection so various session-related variables work.
	// For example: "PRAGMA foreign_keys" for SQLite3, "SET IDENTITY_INSERT" for MS SQL, etc.
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(0)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	// select dialect for driver
	dialect := dialects.ForDriver(driver)
	switch dialect {
	case postgresql.Dialect:
		inspectOnce.Do(func() {
			log.Printf("driver = %q, source = %q", driver, source)

			log.Printf("time.Now()          = %s", now)
			log.Printf("time.Now().UTC()    = %s", now.UTC())

			var version, tz string
			if err = db.QueryRow("SHOW server_version").Scan(&version); err != nil {
				log.Fatal(err)
			}
			if err = db.QueryRow("SHOW TimeZone").Scan(&tz); err != nil {
				log.Fatal(err)
			}
			log.Printf("PostgreSQL version  = %q", version)
			log.Printf("PostgreSQL TimeZone = %q", tz)
		})

	case mysql.Dialect:
		inspectOnce.Do(func() {
			log.Printf("driver = %q, source = %q", driver, source)

			log.Printf("time.Now()          = %s", now)
			log.Printf("time.Now().UTC()    = %s", now.UTC())

			q := "SELECT @@version, @@sql_mode, @@autocommit, @@time_zone"
			var version, mode, autocommit, tz string
			if err = db.QueryRow(q).Scan(&version, &mode, &autocommit, &tz); err != nil {
				log.Fatal(err)
			}
			log.Printf("MySQL version       = %q", version)
			log.Printf("MySQL sql_mode      = %q", mode)
			log.Printf("MySQL autocommit    = %q", autocommit)
			log.Printf("MySQL time_zone     = %q", tz)
		})

	case sqlite3.Dialect:
		if _, err = db.Exec("PRAGMA foreign_keys = ON"); err != nil {
			log.Fatal(err)
		}

		inspectOnce.Do(func() {
			log.Printf("driver = %q, source = %q", driver, source)

			log.Printf("time.Now()          = %s", now)
			log.Printf("time.Now().UTC()    = %s", now.UTC())

			var version, sourceID string
			if err = db.QueryRow("SELECT sqlite_version(), sqlite_source_id()").Scan(&version, &sourceID); err != nil {
				log.Fatal(err)
			}
			log.Printf("SQLite3 version     = %q", version)
			log.Printf("SQLite3 source      = %q", sourceID)
		})

	case mssql.Dialect: //nolint:staticcheck
		fallthrough
	case sqlserver.Dialect:
		inspectOnce.Do(func() {
			log.Printf("driver = %q, source = %q", driver, source)

			log.Printf("time.Now()       = %s", now)
			log.Printf("time.Now().UTC() = %s", now.UTC())

			var version string
			var options uint16
			if err = db.QueryRow("SELECT @@VERSION, @@OPTIONS").Scan(&version, &options); err != nil {
				log.Fatal(err)
			}
			xact := "ON"
			if options&0x4000 == 0 {
				xact = "OFF"
			}
			log.Printf("MS SQL VERSION   = %s", version)
			log.Printf("MS SQL OPTIONS   = %#4x (XACT_ABORT %s)", options, xact)
		})

	default:
		log.Fatalf("reform: no dialect for driver %s", driver)
	}

	return reform.NewDB(db, dialect, nil)
}
