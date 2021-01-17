package gengo

import (
	"go/ast"
	"go/token"
)

func newField(name string, v string, tag string) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: toCamel(name),
			},
		},
		Type: &ast.BasicLit{
			Value: v,
		},
		Tag: &ast.BasicLit{
			Kind:  token.STRING,
			Value: tag,
		},
		Comment: nil,
	}
}

func newStruct(name string, fieldList ast.FieldList) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(toCamel(name)),
				Type: &ast.StructType{
					Fields: &fieldList,
				},
				Comment: nil,
			},
		},
	}
}
