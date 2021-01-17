package gengo

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func Test_newField(t *testing.T) {
	type args struct {
		name string
		v    string
		tag  string
	}
	tests := []struct {
		name string
		args args
		want *ast.Field
	}{
		{name: "NewField", args: args{
			name: "name",
			v:    "value",
			tag:  "tag",
		}, want: &ast.Field{
			Names: []*ast.Ident{
				&ast.Ident{
					Name: toCamel("name"),
				},
			},
			Type: &ast.BasicLit{
				Value: "value",
			},
			Tag: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "tag",
			},
			Comment: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newField(tt.args.name, tt.args.v, tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newStruct(t *testing.T) {
	type args struct {
		name      string
		fieldList ast.FieldList
	}
	list := []*ast.Field{
		newField("name", "value", "tag"),
	}
	tests := []struct {
		name string
		args args
		want *ast.GenDecl
	}{
		{
			name: "NewStruct", args: args{
				name: "name",
				fieldList: ast.FieldList{
					List: list,
				},
			},
			want: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent(toCamel("name")),
						Type: &ast.StructType{
							Fields: &ast.FieldList{List: list},
						},
						Comment: nil,
					},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStruct(tt.args.name, tt.args.fieldList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
