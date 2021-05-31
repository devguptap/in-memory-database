package table

import (
	"razor/column"
)

type Table struct {
	Columns  []column.Column
	DataRows [][]interface{}
}
