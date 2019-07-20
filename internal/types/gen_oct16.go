/*
 This file was autogenerated via
 ---------------------------------------------------------
 gen-builtin --type Oct16 --name oct16 --handler AddUint16
 ---------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Oct16("")

// Oct16 represents field of type oct16
type Oct16 string

// Name returns field name
func (i Oct16) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Oct16) TypeName() string {
	return "oct16"
}

// Register registers a field
func (i Oct16) Register(registrator FieldRegistrator) {
	registrator.AddUint16(i.Name())
}

// Native returns Go's representation of this field's type
func (i Oct16) Native() string {
	return "oct16"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["oct16"] = func(fieldName string) Field {
		return Oct16(fieldName)
	}
}
