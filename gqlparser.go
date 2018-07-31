package gqlparser

import (
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/vektah/gqlparser/parser"
	"github.com/vektah/gqlparser/validator"
	_ "github.com/vektah/gqlparser/validator/rules"
	"github.com/vektah/gqlparser/variable"
)

func LoadSchema(str ...*ast.Source) (*ast.Schema, *gqlerror.Error) {
	return validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
}

func MustLoadSchema(str ...*ast.Source) *ast.Schema {
	s, err := validator.LoadSchema(append([]*ast.Source{validator.Prelude}, str...)...)
	if err != nil {
		panic(err)
	}
	return s
}

func LoadQuery(schema *ast.Schema, str string, coercion variable.CoerceInputScalarFunc) (*ast.QueryDocument, gqlerror.List) {
	query, err := parser.ParseQuery(&ast.Source{Input: str})
	if err != nil {
		return nil, gqlerror.List{err}
	}
	errs := validator.Validate(schema, query, coercion)
	if errs != nil {
		return nil, errs
	}

	return query, nil
}

func MustLoadQuery(schema *ast.Schema, str string, coercion variable.CoerceInputScalarFunc) *ast.QueryDocument {
	q, err := LoadQuery(schema, str, coercion)
	if err != nil {
		panic(err)
	}
	return q
}
