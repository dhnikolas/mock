package gormdb

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	schema.RegisterSerializer("json", JSONSerializer{})
	
	return db, nil
}

type JSONSerializer struct {}

func (JSONSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)
	
	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}
		
		err = json.Unmarshal(bytes, fieldValue.Interface())
	}
	
	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (JSONSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return json.Marshal(fieldValue)
}
