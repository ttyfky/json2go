package descriptor

import (
	"go/ast"
	"io"
	"reflect"
)

// File is source file infomation.
type File struct {
	Name      string
	PkgName   string
	RawFields map[string]interface{}
}

type OutputFile struct {
	Name   string
	Writer io.Writer
	File   *ast.File
}

// Field is Field in json.
type Field struct {
	// Name is Field Name.
	Name string
	// Kind id Field data type.
	Kind reflect.Kind
	// Fields is child Fields.
	Fields []*Field
}
