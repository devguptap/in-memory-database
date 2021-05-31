package app

import (
	"errors"
	"fmt"
	"razor/database"
)

var inMemoryDatabase map[string]*database.Database

func init() {
	inMemoryDatabase = make(map[string]*database.Database)
}

func CreateDatabase(name string) (*database.Database, error) {
	if _, isDBAlreadyExist := inMemoryDatabase[name]; isDBAlreadyExist {
		return nil, errors.New(fmt.Sprintf("Database with name : %s already exist", name))
	} else {
		db := &database.Database{
			Name: name,
		}
		inMemoryDatabase[name] = db
		return db, nil
	}
}

func GetDatabase(name string) (*database.Database, error) {
	if db, ifExist := inMemoryDatabase[name]; ifExist {
		return db, nil
	} else {
		return nil, errors.New(fmt.Sprintf("No database exist with name : %s", name))
	}
}

func DeleteDatabase(name string) error {
	if _, ifExist := inMemoryDatabase[name]; ifExist {
		delete(inMemoryDatabase, name)
		return nil
	} else {
		return errors.New(fmt.Sprintf("No database exist with name : %s", name))
	}
}
