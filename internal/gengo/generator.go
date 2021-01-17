package gengo

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"math"
	"os"
	"reflect"
	"sort"

	"github.com/ttyfky/json2go/descriptor"
	"github.com/ttyfky/json2go/generator"
)

type generationHandle struct {
	reg        *descriptor.Registry
	outputPath string
}

// New create a new generator.Generator.
func New(reg *descriptor.Registry) generator.Generator {
	return &generationHandle{
		reg:        reg,
		outputPath: reg.GetOutputPath(),
	}
}

type outDesc struct {
	file      *ast.File
	path      string
	processed map[string]bool
	writer    io.Writer
}

func (g generationHandle) Generate(files []*descriptor.File) error {
	for _, file := range files {
		o := outDesc{
			file: &ast.File{
				Name:  ast.NewIdent(file.PkgName),
				Decls: make([]ast.Decl, 0, len(file.RawFields)),
			},
			path:      g.outputPath,
			processed: make(map[string]bool),
			writer:    g.reg.GetWriter(),
		}
		o.walk(file.Name, file.RawFields)
		err := o.output(file.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

// walk walks through all the fields in JSON.
// Data types mapping of JSON to Go is described in
// [here](https:// .org/pkg/encoding/json/#Unmarshal).
func (o outDesc) walk(name string, fields map[string]interface{}) {
	list := make([]*ast.Field, 0, len(fields))
	keys := sortedKeys(fields)
	for _, k := range keys {
		fieldDecl := o.mapField(k, fields[k])
		list = append(list, newField(k, fieldDecl, fmt.Sprintf("`json:\"%s,omitempty\"`", k)))
	}
	if _, ok := o.processed[name]; !ok {
		o.processed[name] = true
		o.file.Decls = append(o.file.Decls, newStruct(name,
			ast.FieldList{
				List: list,
			}))
	}
}

func (o outDesc) mapField(k string, v interface{}) string {
	var fieldDecl string

	if v == nil {
		return fmt.Sprintf("%s{}", reflect.Interface.String())
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		fieldDecl = reflect.String.String()
		break
	case reflect.Bool:
		fieldDecl = reflect.Bool.String()
		break
	case reflect.Slice:
		// All items in a list must be the same type.
		elements := v.([]interface{})
		var rawDecl string
		if len(elements) > 0 {
			rawDecl = o.mapField(toSingular(k), elements[0])
		} else {
			rawDecl = fmt.Sprintf("%s{}", reflect.Interface.String())
		}
		fieldDecl = fmt.Sprintf("[]%s", rawDecl)
		break
	case reflect.Map:
		// Create child struct if the type is map.
		o.walk(k, v.(map[string]interface{}))
		fieldDecl = fmt.Sprintf("*%s", toCamel(k))
		break
	case reflect.Float64:
		_, frac := math.Modf(v.(float64))
		if frac > 0 {
			fieldDecl = reflect.Float64.String()
		} else {
			fieldDecl = reflect.Int64.String()
		}
		break
	default:
		fieldDecl = reflect.Invalid.String()
	}
	return fieldDecl
}

func (o outDesc) output(name string) error {
	var w io.Writer
	if o.writer == nil {
		fp, err := os.Create(fmt.Sprintf("%s/%s.go", o.path, name))

		if err != nil {
			return err
		}
		defer fp.Close()
		writer := bufio.NewWriter(fp)
		defer writer.Flush()
		w = writer
	} else {
		w = o.writer
	}

	fmt.Printf("*** generate %s.go ***\n", name)
	err := format.Node(w, token.NewFileSet(), o.file)
	if err != nil {
		return err
	}
	return nil
}

func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
