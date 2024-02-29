// Code generated by github.com/qw4n7y/reform. DO NOT EDIT.

package main

import (
	"fmt"
	"strings"

	"github.com/qw4n7y/reform"
	"github.com/qw4n7y/reform/parse"
)

type tableViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("information_schema").
func (v *tableViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("tables").
func (v *tableViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *tableViewType) Columns() []string {
	return []string{
		"table_catalog",
		"table_schema",
		"table_name",
		"table_type",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *tableViewType) NewStruct() reform.Struct {
	return new(table)
}

// tableView represents tables view or table in SQL database.
var tableView = &tableViewType{
	s: parse.StructInfo{
		Type:      "table",
		SQLSchema: "information_schema",
		SQLName:   "tables",
		Fields: []parse.FieldInfo{
			{Name: "TableCatalog", Type: "string", Column: "table_catalog"},
			{Name: "TableSchema", Type: "string", Column: "table_schema"},
			{Name: "TableName", Type: "string", Column: "table_name"},
			{Name: "TableType", Type: "string", Column: "table_type"},
		},
		PKFieldIndex: -1,
	},
	z: new(table).Values(),
}

// String returns a string representation of this struct or record.
func (s table) String() string {
	res := make([]string, 4)
	res[0] = "TableCatalog: " + reform.Inspect(s.TableCatalog, true)
	res[1] = "TableSchema: " + reform.Inspect(s.TableSchema, true)
	res[2] = "TableName: " + reform.Inspect(s.TableName, true)
	res[3] = "TableType: " + reform.Inspect(s.TableType, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *table) Values() []interface{} {
	return []interface{}{
		s.TableCatalog,
		s.TableSchema,
		s.TableName,
		s.TableType,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *table) Pointers() []interface{} {
	return []interface{}{
		&s.TableCatalog,
		&s.TableSchema,
		&s.TableName,
		&s.TableType,
	}
}

// View returns View object for that struct.
func (s *table) View() reform.View {
	return tableView
}

// check interfaces
var (
	_ reform.View   = tableView
	_ reform.Struct = (*table)(nil)
	_ fmt.Stringer  = (*table)(nil)
)

type columnViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("information_schema").
func (v *columnViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("columns").
func (v *columnViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *columnViewType) Columns() []string {
	return []string{
		"table_catalog",
		"table_schema",
		"table_name",
		"column_name",
		"is_nullable",
		"data_type",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *columnViewType) NewStruct() reform.Struct {
	return new(column)
}

// columnView represents columns view or table in SQL database.
var columnView = &columnViewType{
	s: parse.StructInfo{
		Type:      "column",
		SQLSchema: "information_schema",
		SQLName:   "columns",
		Fields: []parse.FieldInfo{
			{Name: "TableCatalog", Type: "string", Column: "table_catalog"},
			{Name: "TableSchema", Type: "string", Column: "table_schema"},
			{Name: "TableName", Type: "string", Column: "table_name"},
			{Name: "Name", Type: "string", Column: "column_name"},
			{Name: "IsNullable", Type: "yesNo", Column: "is_nullable"},
			{Name: "Type", Type: "string", Column: "data_type"},
		},
		PKFieldIndex: -1,
	},
	z: new(column).Values(),
}

// String returns a string representation of this struct or record.
func (s column) String() string {
	res := make([]string, 6)
	res[0] = "TableCatalog: " + reform.Inspect(s.TableCatalog, true)
	res[1] = "TableSchema: " + reform.Inspect(s.TableSchema, true)
	res[2] = "TableName: " + reform.Inspect(s.TableName, true)
	res[3] = "Name: " + reform.Inspect(s.Name, true)
	res[4] = "IsNullable: " + reform.Inspect(s.IsNullable, true)
	res[5] = "Type: " + reform.Inspect(s.Type, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *column) Values() []interface{} {
	return []interface{}{
		s.TableCatalog,
		s.TableSchema,
		s.TableName,
		s.Name,
		s.IsNullable,
		s.Type,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *column) Pointers() []interface{} {
	return []interface{}{
		&s.TableCatalog,
		&s.TableSchema,
		&s.TableName,
		&s.Name,
		&s.IsNullable,
		&s.Type,
	}
}

// View returns View object for that struct.
func (s *column) View() reform.View {
	return columnView
}

// check interfaces
var (
	_ reform.View   = columnView
	_ reform.Struct = (*column)(nil)
	_ fmt.Stringer  = (*column)(nil)
)

type keyColumnUsageViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("information_schema").
func (v *keyColumnUsageViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("key_column_usage").
func (v *keyColumnUsageViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *keyColumnUsageViewType) Columns() []string {
	return []string{
		"column_name",
		"ordinal_position",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *keyColumnUsageViewType) NewStruct() reform.Struct {
	return new(keyColumnUsage)
}

// keyColumnUsageView represents key_column_usage view or table in SQL database.
var keyColumnUsageView = &keyColumnUsageViewType{
	s: parse.StructInfo{
		Type:      "keyColumnUsage",
		SQLSchema: "information_schema",
		SQLName:   "key_column_usage",
		Fields: []parse.FieldInfo{
			{Name: "ColumnName", Type: "string", Column: "column_name"},
			{Name: "OrdinalPosition", Type: "int", Column: "ordinal_position"},
		},
		PKFieldIndex: -1,
	},
	z: new(keyColumnUsage).Values(),
}

// String returns a string representation of this struct or record.
func (s keyColumnUsage) String() string {
	res := make([]string, 2)
	res[0] = "ColumnName: " + reform.Inspect(s.ColumnName, true)
	res[1] = "OrdinalPosition: " + reform.Inspect(s.OrdinalPosition, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *keyColumnUsage) Values() []interface{} {
	return []interface{}{
		s.ColumnName,
		s.OrdinalPosition,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *keyColumnUsage) Pointers() []interface{} {
	return []interface{}{
		&s.ColumnName,
		&s.OrdinalPosition,
	}
}

// View returns View object for that struct.
func (s *keyColumnUsage) View() reform.View {
	return keyColumnUsageView
}

// check interfaces
var (
	_ reform.View   = keyColumnUsageView
	_ reform.Struct = (*keyColumnUsage)(nil)
	_ fmt.Stringer  = (*keyColumnUsage)(nil)
)

type sqliteMasterViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *sqliteMasterViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("sqlite_master").
func (v *sqliteMasterViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *sqliteMasterViewType) Columns() []string {
	return []string{
		"name",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *sqliteMasterViewType) NewStruct() reform.Struct {
	return new(sqliteMaster)
}

// sqliteMasterView represents sqlite_master view or table in SQL database.
var sqliteMasterView = &sqliteMasterViewType{
	s: parse.StructInfo{
		Type:    "sqliteMaster",
		SQLName: "sqlite_master",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
		},
		PKFieldIndex: -1,
	},
	z: new(sqliteMaster).Values(),
}

// String returns a string representation of this struct or record.
func (s sqliteMaster) String() string {
	res := make([]string, 1)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *sqliteMaster) Values() []interface{} {
	return []interface{}{
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *sqliteMaster) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *sqliteMaster) View() reform.View {
	return sqliteMasterView
}

// check interfaces
var (
	_ reform.View   = sqliteMasterView
	_ reform.Struct = (*sqliteMaster)(nil)
	_ fmt.Stringer  = (*sqliteMaster)(nil)
)

type sqliteTableInfoViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *sqliteTableInfoViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("dummy").
func (v *sqliteTableInfoViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *sqliteTableInfoViewType) Columns() []string {
	return []string{
		"cid",
		"name",
		"type",
		"notnull",
		"dflt_value",
		"pk",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *sqliteTableInfoViewType) NewStruct() reform.Struct {
	return new(sqliteTableInfo)
}

// sqliteTableInfoView represents dummy view or table in SQL database.
var sqliteTableInfoView = &sqliteTableInfoViewType{
	s: parse.StructInfo{
		Type:    "sqliteTableInfo",
		SQLName: "dummy",
		Fields: []parse.FieldInfo{
			{Name: "CID", Type: "int", Column: "cid"},
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Type", Type: "string", Column: "type"},
			{Name: "NotNull", Type: "bool", Column: "notnull"},
			{Name: "DefaultValue", Type: "*string", Column: "dflt_value"},
			{Name: "PK", Type: "int", Column: "pk"},
		},
		PKFieldIndex: -1,
	},
	z: new(sqliteTableInfo).Values(),
}

// String returns a string representation of this struct or record.
func (s sqliteTableInfo) String() string {
	res := make([]string, 6)
	res[0] = "CID: " + reform.Inspect(s.CID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Type: " + reform.Inspect(s.Type, true)
	res[3] = "NotNull: " + reform.Inspect(s.NotNull, true)
	res[4] = "DefaultValue: " + reform.Inspect(s.DefaultValue, true)
	res[5] = "PK: " + reform.Inspect(s.PK, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *sqliteTableInfo) Values() []interface{} {
	return []interface{}{
		s.CID,
		s.Name,
		s.Type,
		s.NotNull,
		s.DefaultValue,
		s.PK,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *sqliteTableInfo) Pointers() []interface{} {
	return []interface{}{
		&s.CID,
		&s.Name,
		&s.Type,
		&s.NotNull,
		&s.DefaultValue,
		&s.PK,
	}
}

// View returns View object for that struct.
func (s *sqliteTableInfo) View() reform.View {
	return sqliteTableInfoView
}

// check interfaces
var (
	_ reform.View   = sqliteTableInfoView
	_ reform.Struct = (*sqliteTableInfo)(nil)
	_ fmt.Stringer  = (*sqliteTableInfo)(nil)
)

func init() {
	parse.AssertUpToDate(&tableView.s, new(table))
	parse.AssertUpToDate(&columnView.s, new(column))
	parse.AssertUpToDate(&keyColumnUsageView.s, new(keyColumnUsage))
	parse.AssertUpToDate(&sqliteMasterView.s, new(sqliteMaster))
	parse.AssertUpToDate(&sqliteTableInfoView.s, new(sqliteTableInfo))
}
