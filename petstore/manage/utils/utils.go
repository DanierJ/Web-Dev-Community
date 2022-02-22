package utils

import (
	"os"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

func DisplayStruct(columns table.Row, rows []table.Row, style table.Style) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(columns)
	t.AppendRows(rows)
	t.SetStyle(style)
	t.Render()
}

func GenerateColumns(s reflect.Value, notDisplayedColumns string) table.Row {
	columns := table.Row{}
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		if strings.Contains(notDisplayedColumns, typeOfT.Field(i).Name) {
			continue
		}
		columns = append(columns, typeOfT.Field(i).Name)
	}

	return columns
}
