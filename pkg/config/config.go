package config

import (
	"fmt"
	"os"
	"reflect"
)

func InitConfiguration(cfg interface{}, verbose bool) error {
	cfgStruct := reflect.ValueOf(cfg).Elem()
	readStruct(cfgStruct, verbose)

	return nil
}

func readStruct(input reflect.Value, verbose bool) {
	inputType := input.Type()

	for i := 0; i < input.NumField(); i++ {
		fieldValue := input.Field(i)
		fieldName := inputType.Field(i).Name

		switch fieldValue.Kind() {
		case reflect.Struct:
			readStruct(fieldValue, verbose)
		case reflect.String:
			setString(fieldName, fieldValue, inputType.Field(i).Tag)
		}

		if verbose && fieldValue.Kind() != reflect.Struct {
			fmt.Printf("%s: %v\n", fieldName, fieldValue)
		}

	}

}

func setString(name string, fieldValue reflect.Value, tag reflect.StructTag) {
	envTag := tag.Get("env")
	value := os.Getenv(envTag)

	if value == "" {
		value = tag.Get("default")
	}

	// Ensure the value is addressable
	if fieldValue.CanSet() {
		// Set the field value
		fieldValue.SetString(value)
	}
}
