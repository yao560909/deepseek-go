package apijson

import (
	"github.com/tidwall/gjson"
	"reflect"
)

var unionVariants = map[reflect.Type]interface{}{}

type UnionVariant struct {
	TypeFilter         gjson.Type
	DiscriminatorValue interface{}
	Type               reflect.Type
}

var unionRegistry = map[reflect.Type]unionEntry{}

type unionEntry struct {
	discriminatorKey string
	variants         []UnionVariant
}
