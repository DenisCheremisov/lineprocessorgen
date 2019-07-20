/*
 This file was autogenerated via
 ---------------------------------------------------------
 gen-builtin --declarable --native --type Int8 --name int8
 ---------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Int8("")

// Int8 represents field of type int8
type Int8 string

// Name returns field name
func (i Int8) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Int8) TypeName() string {
	return "int8"
}

// Register registers a field
func (i Int8) Register(registrator FieldRegistrator) {
	registrator.AddInt8(i.Name())
}

// GoName returns Go's representation of this field's type
func (i Int8) GoName() string {
	return "int8"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["int8"] = func(fieldName string) Field {
		return Int8(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["int8"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["int8"] = struct{}{}

}
