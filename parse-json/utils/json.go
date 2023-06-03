package utils

import (
	"reflect"
	"strings"
)

type JsonObjectParams interface {
	Keys(data interface{}) []string
	Values(data interface{}) []interface{}
	Entries(data interface{}) []map[string]interface{}
}

type JsonObject struct{}

func (o JsonObject) Keys(data interface{}) []string {

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil
	}

	keys := make([]string, v.NumField())
	typeOfData := v.Type()
	for i := 0; i < v.NumField(); i++ {
		keys[i] = getJsonField(string(typeOfData.Field(i).Tag))
	}

	return keys
}

func (o JsonObject) Values(data interface{}) []interface{} {

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil
	}

	keys := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {

		keys[i] = v.Field(i).Interface()
	}

	return keys
}

func (o JsonObject) Entries(data interface{}) [][]interface{} {

	ks, vs := o.Keys(data), o.Values(data)

	entires := make([][]interface{}, len(ks))
	for i, v := range ks {
		entires[i] = []interface{}{v, vs[i]}
	}

	return entires
}

func getJsonField(jsonStr string) string {

	fields := strings.Split(jsonStr, "json:")[1]
	return strings.Trim(fields, "\"")
}
