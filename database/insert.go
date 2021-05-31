package database

import (
	"errors"
	"fmt"
)

func (d *Database) Insert(tableName string, data []interface{}) error {
	if table, ok := d.Tables[tableName]; !ok {
		return errors.New(fmt.Sprintf("Table with name : %s not exist", tableName))
	} else {
		lengthOfColumn := len(table.Columns)
		if lengthOfColumn != len(data) {
			return errors.New(fmt.Sprintf("Length of data : %s is not matching with table schema", len(data)))
		}

		var row = make([]interface{}, lengthOfColumn)
		var err error
		for i := 0; i < lengthOfColumn; i++ {
			if err = table.Columns[i].DataType.ValidateValue(data[i]); err != nil {
				return err
			}

			if err = table.Columns[i].ValidateValueAgainstConstraints(data[i]); err != nil {
				return err
			}
			row[i] = data[i]
		}

		table.DataRows = append(table.DataRows, row)
		return nil
	}
}
