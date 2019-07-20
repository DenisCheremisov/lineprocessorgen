/*
 This file was autogenerated via
 ---------------------------------
 gen-builtin --type Str --name str
 ---------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Str("")

// Str represents field of type str
type Str string

// Name returns field name
func (i Str) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Str) TypeName() string {
	return "str"
}

// Register registers a field
func (i Str) Register(registrator FieldRegistrator) {
	registrator.AddStr(i.Name())
}

// Native returns Go's representation of this field's type
func (i Str) Native() string {
	return "str"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["str"] = func(fieldName string) Field {
		return Str(fieldName)
	}
}
