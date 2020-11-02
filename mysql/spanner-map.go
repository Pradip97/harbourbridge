package mysql

import (
	"github.com/cloudspannerecosystem/harbourbridge/internal"
	"github.com/cloudspannerecosystem/harbourbridge/spanner/ddl"
)

var ToSpannerType = map[string]*struct {
	MaxMods int
	T       ddl.Type
	Issue   []internal.SchemaIssue
}{
	"bool":       {MaxMods: 0, T: ddl.Type{Name: ddl.Bool, Len: ddl.MaxLength}, Issue: nil},
	"varchar":    {MaxMods: 0, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"char":       {MaxMods: 0, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"text":       {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"tinytext":   {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"mediumtext": {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"longtext":   {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"set":        {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"enum":       {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"json":       {MaxMods: 0, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: nil},
	"bit":        {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"binary":     {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"varbinary":  {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"blob":       {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"tinyblob":   {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"mediumblob": {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"longblob":   {MaxMods: 1, T: ddl.Type{Name: ddl.Bytes, Len: ddl.MaxLength}, Issue: nil},
	"tinyint":    {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"smallint":   {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"mediumint":  {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"int":        {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"integer":    {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"bigint":     {MaxMods: 1, T: ddl.Type{Name: ddl.Int64}, Issue: nil},
	"double":     {MaxMods: 2, T: ddl.Type{Name: ddl.Float64}, Issue: nil},
	"float":      {MaxMods: 2, T: ddl.Type{Name: ddl.Float64}, Issue: []internal.SchemaIssue{internal.Widened}},
	"numeric":    {MaxMods: 2, T: ddl.Type{Name: ddl.Float64}, Issue: []internal.SchemaIssue{internal.Decimal}},
	"date":       {MaxMods: 1, T: ddl.Type{Name: ddl.Date}, Issue: nil},
	"datetime":   {MaxMods: 1, T: ddl.Type{Name: ddl.Timestamp}, Issue: []internal.SchemaIssue{internal.Datetime}},
	"timestamp":  {MaxMods: 1, T: ddl.Type{Name: ddl.Timestamp}, Issue: nil},
	"time":       {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: []internal.SchemaIssue{internal.Time}},
	"year":       {MaxMods: 1, T: ddl.Type{Name: ddl.String, Len: ddl.MaxLength}, Issue: []internal.SchemaIssue{internal.Time}},
}
