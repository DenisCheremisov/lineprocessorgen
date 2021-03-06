/*
 This file was autogenerated via
 ---------------------------------------------------------
 gen-builtin --declarable --native --type Uint --name uint
 ---------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Uint("")

// Uint represents field of type uint
type Uint string

// Name returns field name
func (i Uint) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Uint) TypeName() string {
	return "uint"
}

// Register registers a field
func (i Uint) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddUint(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Uint) GoName() string {
	return "uint"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["uint"] = func(fieldName string) Field {
		return Uint(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["uint"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["uint"] = struct{}{}

}
