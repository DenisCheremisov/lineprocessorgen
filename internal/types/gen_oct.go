/*
 This file was autogenerated via
 ---------------------------------------------------
 gen-builtin --type Oct --name oct --handler AddUint
 ---------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Oct("")

// Oct represents field of type oct
type Oct string

// Name returns field name
func (i Oct) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Oct) TypeName() string {
	return "oct"
}

// Register registers a field
func (i Oct) Register(registrator FieldRegistrator) {
	registrator.AddUint(i.Name())
}

// Native returns Go's representation of this field's type
func (i Oct) Native() string {
	return "oct"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["oct"] = func(fieldName string) Field {
		return Oct(fieldName)
	}
}
