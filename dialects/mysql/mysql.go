// Package mysql implements reform.Dialect for MySQL.
package mysql // import "github.com/qw4n7y/reform/dialects/mysql"

import (
	"github.com/qw4n7y/reform"
)

type mysql struct{}

func (mysql) String() string {
	return "mysql"
}

func (mysql) Placeholder(index int) string {
	return "?"
}

func (mysql) Placeholders(start, count int) []string {
	res := make([]string, count)
	for i := 0; i < count; i++ {
		res[i] = "?"
	}
	return res
}

func (mysql) QuoteIdentifier(identifier string) string {
	return "`" + identifier + "`"
}

func (mysql) LastInsertIdMethod() reform.LastInsertIdMethod {
	return reform.LastInsertId
}

func (mysql) SelectLimitMethod() reform.SelectLimitMethod {
	return reform.Limit
}

func (mysql) DefaultValuesMethod() reform.DefaultValuesMethod {
	return reform.EmptyLists
}

// Dialect implements reform.Dialect for MySQL.
var Dialect mysql

// check interface
var _ reform.Dialect = Dialect
