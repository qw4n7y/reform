// Package mssql implements reform.Dialect for Microsoft SQL Server (mssql driver).
//
// Deprecated: Use sqlserver dialect instead. https://github.com/denisenkom/go-mssqldb#deprecated
package mssql // import "github.com/qw4n7y/reform/dialects/mssql"

import (
	"github.com/qw4n7y/reform"
)

type mssql struct{}

func (mssql) String() string {
	return "mssql"
}

func (mssql) Placeholder(index int) string {
	return "?"
}

func (mssql) Placeholders(start, count int) []string {
	res := make([]string, count)
	for i := 0; i < count; i++ {
		res[i] = "?"
	}
	return res
}

func (mssql) QuoteIdentifier(identifier string) string {
	return "[" + identifier + "]"
}

func (mssql) LastInsertIdMethod() reform.LastInsertIdMethod {
	return reform.OutputInserted
}

func (mssql) SelectLimitMethod() reform.SelectLimitMethod {
	return reform.SelectTop
}

func (mssql) DefaultValuesMethod() reform.DefaultValuesMethod {
	return reform.DefaultValues
}

// Dialect implements reform.Dialect for Microsoft SQL Server.
//
// Deprecated: Use sqlserver.Dialect instead. https://github.com/denisenkom/go-mssqldb#deprecated
var Dialect mssql

// check interface
var _ reform.Dialect = Dialect
