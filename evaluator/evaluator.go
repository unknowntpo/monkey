package evaluator

import (
	"github.com/unknowntpo/monkey/ast"
	"github.com/unknowntpo/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}
