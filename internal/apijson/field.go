package apijson

import "reflect"

type status uint8

const (
	missing status = iota
	null
	invalid
	valid
)

type Field struct {
	raw    string
	status status
}

func (j Field) IsNull() bool    { return j.status <= null }
func (j Field) IsMissing() bool { return j.status == missing }
func (j Field) IsInvalid() bool { return j.status == invalid }
func (j Field) Raw() string     { return j.raw }

func getSubField(root reflect.Value, index []int, name string) reflect.Value {
	strct := root.FieldByIndex(index[:len(index)-1])
	if !strct.IsValid() {
		panic("couldn't find encapsulating struct for field " + name)
	}
	meta := strct.FieldByName("JSON")
	if !meta.IsValid() {
		return reflect.Value{}
	}
	field := meta.FieldByName(name)
	if !field.IsValid() {
		return reflect.Value{}
	}
	return field
}
