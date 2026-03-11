package utils

import (
	"go/ast"
	"go/token"
	"strings"
)

func IsLogCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	method := sel.Sel.Name

	switch method {

	case "Info", "Error", "Warn", "Debug":
		return true
	}

	return false
}

func GetMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return "", false
	}

	if lit.Kind != token.STRING {
		return "", false
	}

	msg := strings.Trim(lit.Value, `"`)

	return msg, true
}
