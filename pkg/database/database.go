package database

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"github.com/weblair/lair/pkg/model"
)

type NotFoundError struct {
	FieldName string
	FieldValue interface{}
	ResourceType reflect.Type
}

func (nfe NotFoundError) Error() string {
	return fmt.Sprintf("%v record with %s = %v not found", nfe.ResourceType, nfe.FieldName, nfe.FieldValue)
}

type DBI interface {
	Fetch(param string, value interface{}, recordType reflect.Type, preload []string)
}

type Database struct {
	DB *gorm.DB
}

func (db Database) List(model *[]model.Model) error {
	db.DB.Find(model)

	return nil
}

func (db Database) Fetch(model *model.Model, pName string, pValue interface{}, preload []string) error {
	tx := db.DB.Where(fmt.Sprintf("%s = ?", pName), pValue)

	for _, v := range preload {
		tx = tx.Preload(v)
	}

	tx.First(model)

	if (*model).GetID() == 0 {
		return NotFoundError{
			ResourceType: reflect.TypeOf(*model),
			FieldName: pName,
			FieldValue: pValue,
		}
	}

	return nil
}

var DB Database
