package database

import (
	"errors"
	"fmt"
	"razor/column"
	"razor/constants"
	"razor/table"
)

type Database struct {
	Name   string
	Tables map[string]*table.Table
}

func (d *Database) CreateTable(tableName string, columnNames, columnDataType []string, constraints [][]string) error {
	if tableName == "" {
		return errors.New("table name should not be empty")
	}

	if len(d.Tables) != 0 {
		if _, exist := d.Tables[tableName]; exist {
			return errors.New(fmt.Sprintf("table name : %s should not be empty", tableName))
		}
	} else {
		d.Tables = make(map[string]*table.Table)
	}

	if err := d.runAllColumnDataValidations(columnNames, columnDataType, constraints); err != nil {
		return err
	}

	t := &table.Table{
		Columns:  make([]column.Column, 0),
		DataRows: make([][]interface{}, 0),
	}
	for i := 0; i < len(columnNames); i++ {
		t.Columns = append(t.Columns, column.NewColumn(columnNames[i], columnDataType[i], constraints[i]))
	}

	d.Tables[tableName] = t
	return nil

}

func (d *Database) runAllColumnDataValidations(columnNames, columnDataType []string, constraints [][]string) error {
	if !d.isValidColumnNames(columnNames) {
		return errors.New("column name should not be empty")
	}

	if !d.isSameLength(columnNames, columnDataType, constraints) {
		return errors.New("length of columnNames,columnDataType and  constraints should be same")
	}

	if !d.isValidColumnDataType(columnDataType) {
		return errors.New(fmt.Sprintf("invalid column data type received. Supported types are %s, %s", constants.StringDataType, constants.IntegerDataType))
	}

	if !d.isValidConstraintType(constraints) {
		return errors.New(fmt.Sprintf("invalid constraint type received. Supported type is %s", constants.RequiredConstraint))
	}
	return nil
}

func (d *Database) isValidColumnDataType(columnDataType []string) bool {
	for _, dataType := range columnDataType {
		if !(dataType == constants.StringDataType || dataType == constants.IntegerDataType) {
			return false
		}
	}
	return true
}

func (d *Database) isValidConstraintType(constraintsList [][]string) bool {
	for _, constraints := range constraintsList {
		for _, constraint := range constraints {
			if constraint != "" && constraint != constants.RequiredConstraint {
				return false
			}
		}

	}
	return true
}

func (d *Database) isSameLength(columnNames, columnDataType []string, constraints [][]string) bool {
	if len(columnNames) == len(columnDataType) && len(columnNames) == len(constraints) {
		return true
	}
	return false
}

func (d *Database) isValidColumnNames(columnNames []string) bool {
	for _, name := range columnNames {
		if name == "" {
			return false
		}
	}
	return true
}
