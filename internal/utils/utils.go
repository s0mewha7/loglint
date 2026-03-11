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

	switch sel.Sel.Name {
	case "Info", "Error", "Warn", "Debug", "Fatal":
		return true
	}

	return false
}

func GetMessage(call *ast.CallExpr) (string, bool) {
	if len(call.Args) == 0 {
		return "", false
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", false
	}

	return strings.Trim(lit.Value, `"`), true
}
